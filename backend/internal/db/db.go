package db

import (
	"backend/internal/config"
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init(cfg *config.Config) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, cfg.DB_URL)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping Postgres: %v", err)
	}

	log.Println("Connected to Postgres successfully")
	Pool = pool

	createTables(ctx)

	return Pool
}

func Close() {
	if Pool != nil {
		Pool.Close()
		log.Println("Postgres pool closed")
	}
}

func createTables(ctx context.Context) {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id BIGSERIAL PRIMARY KEY,
			verified BOOLEAN DEFAULT FALSE,
			role VARCHAR(16) NOT NULL DEFAULT 'client'
				CHECK (role IN ('client', 'server', 'superadmin')),
			status VARCHAR(16) NOT NULL DEFAULT 'review'
				CHECK (status IN ('active', 'review', 'suspended', 'banned')),
			name VARCHAR(32),
			avatar VARCHAR(512),
			bio VARCHAR(512),
			phone VARCHAR(24) UNIQUE,
			email VARCHAR(64) UNIQUE,
			password VARCHAR(512),
			reset_token VARCHAR(256),
			reset_token_expiry TIMESTAMPTZ,
			google_id VARCHAR(128) UNIQUE,
			google_id_token VARCHAR(128),
			google_access_token VARCHAR(128),
			fcm_token VARCHAR(128),
			refresh_token VARCHAR(128),
			refresh_token_at TIMESTAMPTZ,
			rating_avg NUMERIC(3,2) NOT NULL DEFAULT 0.00,
			rating_count INT NOT NULL DEFAULT 0,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE services (
			id BIGSERIAL PRIMARY KEY,
			active BOOLEAN NOT NULL DEFAULT TRUE,
			user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			category_id BIGINT NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,
			subcategory_id BIGINT NOT NULL REFERENCES sub_categories(id) ON DELETE RESTRICT,
			division_id INT NOT NULL,
			district_id INT NOT NULL,
			subdistrict_id INT NOT NULL,
			area VARCHAR(256) NOT NULL,
			title VARCHAR(64) NOT NULL,
			caption VARCHAR(256) NOT NULL,
			description VARCHAR(1024) NOT NULL,
			price VARCHAR(32) NOT NULL,
			features JSONB CHECK (features IS NULL OR (jsonb_typeof(features) = 'object' AND length(features::text) <= 4096)),
			hours VARCHAR(48) CHECK (hours IS NULL OR hours = 'All day' OR hours ~ '^([01]?[0-9]|2[0-3]):[0-5][0-9]-([01]?[0-9]|2[0-3]):[0-5][0-9]$'),
			days TEXT[] NOT NULL CHECK (ARRAY(SELECT unnest(days) EXCEPT SELECT unnest(ARRAY['mon','tue','wed','thu','fri','sat','sun'])) = '{}' AND length(array_to_string(days,',')) <= 32),
			page_name VARCHAR(32),
			page_link VARCHAR(256),
			messenger_name VARCHAR(32),
			messenger_link VARCHAR(256),
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);`,
	}

	for _, q := range tables {
		if _, err := Pool.Exec(ctx, q); err != nil {
			log.Fatalf("Failed to create table: %v", err)
		}
	}

	log.Println("All tables ensured")

	indexes := []string{
		// Users indexes
		`CREATE UNIQUE INDEX IF NOT EXISTS idx_users_google_id ON users(google_id);`,
		`CREATE INDEX IF NOT EXISTS idx_users_phone ON users(phone);`,
		`CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);`,
		`CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);`,
		`CREATE INDEX IF NOT EXISTS idx_users_refresh_token ON users(refresh_token);`,

		// Services indexes
		`CREATE INDEX IF NOT EXISTS idx_services_user_id ON services(user_id);`,
		`CREATE INDEX IF NOT EXISTS idx_services_active ON services(active);`,
		`CREATE INDEX IF NOT EXISTS idx_services_category_id ON services(category_id);`,
		`CREATE INDEX IF NOT EXISTS idx_services_subcategory_id ON services(subcategory_id);`,
		`CREATE INDEX IF NOT EXISTS idx_services_created_at ON services(created_at);`,
		`CREATE INDEX IF NOT EXISTS idx_services_rating ON services(rating_avg);`,
		`CREATE INDEX IF NOT EXISTS idx_services_category_location_active 
			ON services(category_id, division_id, district_id, subdistrict_id) 
			WHERE active = TRUE;`,
		`CREATE INDEX IF NOT EXISTS idx_services_location ON services(division_id, district_id, subdistrict_id);`,
		`CREATE INDEX IF NOT EXISTS idx_services_features ON services USING GIN(features);`,
		`CREATE INDEX IF NOT EXISTS idx_services_days ON services USING GIN(days);`,

		// Service bookings indexes
		`CREATE INDEX IF NOT EXISTS idx_bookings_user_id ON service_bookings(user_id);`,
		`CREATE INDEX IF NOT EXISTS idx_bookings_provider_id ON service_bookings(provider_id);`,
		`CREATE INDEX IF NOT EXISTS idx_bookings_service_id ON service_bookings(service_id);`,
		`CREATE INDEX IF NOT EXISTS idx_bookings_status ON service_bookings(status);`,
		`CREATE INDEX IF NOT EXISTS idx_bookings_time ON service_bookings(start_time, end_time);`,
	}

	for _, i := range indexes {
		if _, err := Pool.Exec(ctx, i); err != nil {
			log.Fatalf("Failed to create index: %v", err)
		}
	}

	log.Println("All indexes ensured")
}

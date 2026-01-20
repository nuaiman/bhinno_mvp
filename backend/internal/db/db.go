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
			role VARCHAR(25) NOT NULL DEFAULT 'client'
				CHECK (role IN ('client', 'server', 'superadmin')),
			status VARCHAR(25) NOT NULL DEFAULT 'review'
				CHECK (status IN ('active', 'review', 'suspended', 'banned')),
			name VARCHAR(255),
			phone VARCHAR(25) UNIQUE NOT NULL,
			password VARCHAR(1000) NOT NULL,
			email VARCHAR(255) UNIQUE,
			avatar VARCHAR(255),
			bio VARCHAR(255),
			social_media_1_name VARCHAR(25),
			social_media_1_url VARCHAR(255),
			social_media_2_name VARCHAR(25),
			social_media_2_url VARCHAR(255),
			rating_avg NUMERIC(3,2) DEFAULT 0,
			rating_count INT DEFAULT 0,
			refresh_token VARCHAR(255),
			refresh_token_at TIMESTAMP,
			created_at TIMESTAMP DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS categories (
			id BIGSERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS sub_categories (
			id BIGSERIAL PRIMARY KEY,
			category_id BIGINT NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
			name VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		);`,
	}

	for _, q := range tables {
		if _, err := Pool.Exec(ctx, q); err != nil {
			log.Fatalf("Failed to create table: %v", err)
		}
	}

	log.Println("All tables ensured")

	indexes := []string{
		`CREATE INDEX IF NOT EXISTS idx_users_phone ON users(phone);`,
		`CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);`,
		`CREATE INDEX IF NOT EXISTS idx_profiles_user_id ON profiles(user_id);`,
		`CREATE INDEX IF NOT EXISTS idx_profiles_status ON profiles(status);`,
		`CREATE INDEX IF NOT EXISTS idx_profiles_rating ON profiles(rating_avg);`,
		`CREATE INDEX IF NOT EXISTS idx_sub_categories_category_id ON sub_categories(category_id);`,
	}

	for _, i := range indexes {
		if _, err := Pool.Exec(ctx, i); err != nil {
			log.Fatalf("Failed to create index: %v", err)
		}
	}

	log.Println("All indexes ensured")
}

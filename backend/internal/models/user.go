package models

import (
	"backend/internal/db"
	"context"
	"log"
	"time"
)

type User struct {
	ID                int64      `json:"id"`
	Name              string     `json:"name,omitempty"`
	Phone             string     `json:"phone,omitempty"`
	Email             string     `json:"email,omitempty"`
	Password          string     `json:"-"`
	Verified          bool       `json:"verified,omitempty"`
	Role              string     `json:"role,omitempty"`
	Status            string     `json:"status,omitempty"`
	Avatar            string     `json:"avatar,omitempty"`
	Bio               string     `json:"bio,omitempty"`
	ResetToken        string     `json:"-"`
	ResetTokenExpiry  *time.Time `json:"-"`
	GoogleID          string     `json:"-"`
	GoogleIDToken     string     `json:"-"`
	GoogleAccessToken string     `json:"-"`
	FCMToken          string     `json:"-"`
	RefreshToken      string     `json:"-"`
	RefreshTokenAt    *time.Time `json:"-"`
	RatingAvg         float64    `json:"rating_avg,omitempty"`
	RatingCount       int        `json:"rating_count,omitempty"`
	CreatedAt         time.Time  `json:"created_at,omitzero"`
}

func EnsureSuperAdmin(email, hashedPassword string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id int64
	err := db.Pool.QueryRow(ctx, `SELECT id FROM users WHERE role='superadmin'`).Scan(&id)
	if err != nil {
		_, err := db.Pool.Exec(ctx, `
			INSERT INTO users (email, password, role)
			VALUES ($1, $2, 'superadmin')
		`, email, hashedPassword)
		if err != nil {
			log.Fatalf("Failed to create superadmin: %v", err)
		}
		log.Println("Superadmin created")
		return
	}

	_, err = db.Pool.Exec(ctx, `
		UPDATE users
		SET email=$1, password=$2
		WHERE id=$3
	`, email, hashedPassword, id)
	if err != nil {
		log.Fatalf("Failed to update superadmin: %v", err)
	}
	log.Println("Superadmin updated")
}

func CreateUserWithGoogle(ctx context.Context, u *User) error {
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO users (email, google_id, name, avatar, verified)
		VALUES ($1, $2, $3, $4, $5)
	`, u.Email, u.GoogleID, u.Name, u.Avatar, u.Verified)
	return err
}

func CreateUserWithEmail(ctx context.Context, u *User) error {
	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
	`
	_, err := db.Pool.Exec(ctx, query, u.Email, u.Password)
	return err
}

func GetUserByID(ctx context.Context, userID int64) (*User, error) {
	u := &User{}
	err := db.Pool.QueryRow(ctx, `
		SELECT
			id, verified, role, status, name, avatar, bio,
			phone, email, password,
			reset_token, reset_token_expiry,
			google_id, google_id_token, google_access_token,
			fcm_token, refresh_token, refresh_token_at,
			rating_avg, rating_count, created_at
		FROM users
		WHERE id=$1
	`, userID).Scan(
		&u.ID, &u.Verified, &u.Role, &u.Status, &u.Name, &u.Avatar, &u.Bio,
		&u.Phone, &u.Email, &u.Password,
		&u.ResetToken, &u.ResetTokenExpiry,
		&u.GoogleID, &u.GoogleIDToken, &u.GoogleAccessToken,
		&u.FCMToken, &u.RefreshToken, &u.RefreshTokenAt,
		&u.RatingAvg, &u.RatingCount, &u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByGoogleID(ctx context.Context, googleID string) (*User, error) {
	u := &User{}
	err := db.Pool.QueryRow(ctx, `
		SELECT
			id, verified, role, status, name, avatar, bio, phone, email, password,
			reset_token, reset_token_expiry, google_id, google_id_token, google_access_token,
			fcm_token, refresh_token, refresh_token_at, rating_avg, rating_count, created_at
		FROM users
		WHERE google_id=$1
	`, googleID).Scan(
		&u.ID, &u.Verified, &u.Role, &u.Status, &u.Name, &u.Avatar, &u.Bio, &u.Phone, &u.Email, &u.Password,
		&u.ResetToken, &u.ResetTokenExpiry, &u.GoogleID, &u.GoogleIDToken, &u.GoogleAccessToken,
		&u.FCMToken, &u.RefreshToken, &u.RefreshTokenAt, &u.RatingAvg, &u.RatingCount, &u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByPhone(ctx context.Context, phone string) (*User, error) {
	u := &User{}
	err := db.Pool.QueryRow(ctx, `
		SELECT
			id, verified, role, status, name, avatar, bio,
			phone, email, password,
			reset_token, reset_token_expiry,
			google_id, google_id_token, google_access_token,
			fcm_token, refresh_token, refresh_token_at,
			rating_avg, rating_count, created_at
		FROM users
		WHERE phone=$1
	`, phone).Scan(
		&u.ID, &u.Verified, &u.Role, &u.Status, &u.Name, &u.Avatar, &u.Bio,
		&u.Phone, &u.Email, &u.Password,
		&u.ResetToken, &u.ResetTokenExpiry,
		&u.GoogleID, &u.GoogleIDToken, &u.GoogleAccessToken,
		&u.FCMToken, &u.RefreshToken, &u.RefreshTokenAt,
		&u.RatingAvg, &u.RatingCount, &u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := &User{}
	err := db.Pool.QueryRow(ctx, `
		SELECT
			id, verified, role, status, name, avatar, bio,
			phone, email, password,
			reset_token, reset_token_expiry,
			google_id, google_id_token, google_access_token,
			fcm_token, refresh_token, refresh_token_at,
			rating_avg, rating_count, created_at
		FROM users
		WHERE email=$1
	`, email).Scan(
		&u.ID, &u.Verified, &u.Role, &u.Status, &u.Name, &u.Avatar, &u.Bio,
		&u.Phone, &u.Email, &u.Password,
		&u.ResetToken, &u.ResetTokenExpiry,
		&u.GoogleID, &u.GoogleIDToken, &u.GoogleAccessToken,
		&u.FCMToken, &u.RefreshToken, &u.RefreshTokenAt,
		&u.RatingAvg, &u.RatingCount, &u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByRefreshToken(ctx context.Context, token string) (*User, error) {
	u := &User{}
	err := db.Pool.QueryRow(ctx, `
		SELECT
			id, verified, role, status, name, avatar, bio,
			phone, email, password,
			reset_token, reset_token_expiry,
			google_id, google_id_token, google_access_token,
			fcm_token, refresh_token, refresh_token_at,
			rating_avg, rating_count, created_at
		FROM users
		WHERE refresh_token=$1
	`, token).Scan(
		&u.ID, &u.Verified, &u.Role, &u.Status, &u.Name, &u.Avatar, &u.Bio,
		&u.Phone, &u.Email, &u.Password,
		&u.ResetToken, &u.ResetTokenExpiry,
		&u.GoogleID, &u.GoogleIDToken, &u.GoogleAccessToken,
		&u.FCMToken, &u.RefreshToken, &u.RefreshTokenAt,
		&u.RatingAvg, &u.RatingCount, &u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UpdateUserRefreshToken(ctx context.Context, userID int64, token string) error {
	var err error
	if token == "" {
		_, err = db.Pool.Exec(ctx, `
			UPDATE users
			SET refresh_token = NULL, refresh_token_at = NULL
			WHERE id = $1
		`, userID)
	} else {
		_, err = db.Pool.Exec(ctx, `
			UPDATE users
			SET refresh_token = $1, refresh_token_at = NOW()
			WHERE id = $2
		`, token, userID)
	}
	return err
}

func UpdateUser(ctx context.Context, u *User) error {
	_, err := db.Pool.Exec(ctx, `
		UPDATE users
		SET phone=$1, name=$2, email=$3, avatar=$4, bio=$5,
		    verified=$6, status=$7
		WHERE id=$8
	`, u.Phone, u.Name, u.Email, u.Avatar, u.Bio,
		u.Verified, u.Status, u.ID)
	return err
}

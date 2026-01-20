package models

import (
	"backend/internal/db"
	"context"
	"log"
	"time"
)

type User struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name,omitempty"`
	Phone            string    `json:"phone,omitempty"`
	Password         string    `json:"-"` // hide password in JSON
	RefreshToken     string    `json:"-"`
	RefreshTokenAt   time.Time `json:"-"`
	Role             string    `json:"role,omitempty"`
	Status           string    `json:"status,omitempty"`
	Verified         bool      `json:"verified,omitempty"`
	Email            string    `json:"email,omitempty"`
	Avatar           string    `json:"avatar,omitempty"`
	Bio              string    `json:"bio,omitempty"`
	SocialMedia1Name string    `json:"social_media_1_name,omitempty"`
	SocialMedia1URL  string    `json:"social_media_1_url,omitempty"`
	SocialMedia2Name string    `json:"social_media_2_name,omitempty"`
	SocialMedia2URL  string    `json:"social_media_2_url,omitempty"`
	RatingAvg        float64   `json:"rating_avg,omitempty"`
	RatingCount      int       `json:"rating_count,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
}

// EnsureSuperAdmin creates or updates the superadmin user
func EnsureSuperAdmin(phone, hashedPassword string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id int64
	err := db.Pool.QueryRow(ctx, `SELECT id FROM users WHERE role='superadmin'`).Scan(&id)
	if err != nil {
		_, err := db.Pool.Exec(ctx, `
			INSERT INTO users (phone, password, role)
			VALUES ($1, $2, 'superadmin')
		`, phone, hashedPassword)
		if err != nil {
			log.Fatalf("Failed to create superadmin: %v", err)
		}
		log.Println("Superadmin created")
		return
	}

	_, err = db.Pool.Exec(ctx, `
		UPDATE users
		SET phone=$1, password=$2
		WHERE id=$3
	`, phone, hashedPassword, id)
	if err != nil {
		log.Fatalf("Failed to update superadmin: %v", err)
	}
	log.Println("Superadmin updated")
}

// Fetch user by ID
func GetUserByID(ctx context.Context, userID int64) (*User, error) {
	u := &User{}
	err := db.Pool.QueryRow(ctx, `
		SELECT 
			id, phone, password, refresh_token, refresh_token_at, role, status, name, email, avatar, bio,
			social_media_1_name, social_media_1_url, social_media_2_name, social_media_2_url,
			rating_avg, rating_count, verified, created_at
		FROM users
		WHERE id=$1
	`, userID).Scan(
		&u.ID, &u.Phone, &u.Password, &u.RefreshToken, &u.RefreshTokenAt,
		&u.Role, &u.Status, &u.Name, &u.Email, &u.Avatar, &u.Bio,
		&u.SocialMedia1Name, &u.SocialMedia1URL, &u.SocialMedia2Name, &u.SocialMedia2URL,
		&u.RatingAvg, &u.RatingCount, &u.Verified, &u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Fetch user by phone
func GetUserByPhone(ctx context.Context, phone string) (*User, error) {
	u := &User{}
	err := db.Pool.QueryRow(ctx, `
		SELECT 
			id, phone, password, refresh_token, refresh_token_at, role, status, name, email, avatar, bio,
			social_media_1_name, social_media_1_url, social_media_2_name, social_media_2_url,
			rating_avg, rating_count, verified, created_at
		FROM users
		WHERE phone=$1
	`, phone).Scan(
		&u.ID, &u.Phone, &u.Password, &u.RefreshToken, &u.RefreshTokenAt,
		&u.Role, &u.Status, &u.Name, &u.Email, &u.Avatar, &u.Bio,
		&u.SocialMedia1Name, &u.SocialMedia1URL, &u.SocialMedia2Name, &u.SocialMedia2URL,
		&u.RatingAvg, &u.RatingCount, &u.Verified, &u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Fetch user by Email
func GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := &User{}
	err := db.Pool.QueryRow(ctx, `
		SELECT 
			id, phone, password, refresh_token, refresh_token_at, role, status, name, email, avatar, bio,
			social_media_1_name, social_media_1_url, social_media_2_name, social_media_2_url,
			rating_avg, rating_count, verified, created_at
		FROM users
		WHERE email=$1
	`, email).Scan(
		&u.ID, &u.Phone, &u.Password, &u.RefreshToken, &u.RefreshTokenAt,
		&u.Role, &u.Status, &u.Name, &u.Email, &u.Avatar, &u.Bio,
		&u.SocialMedia1Name, &u.SocialMedia1URL, &u.SocialMedia2Name, &u.SocialMedia2URL,
		&u.RatingAvg, &u.RatingCount, &u.Verified, &u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Fetch user by refresh token
func GetUserByRefreshToken(ctx context.Context, token string) (*User, error) {
	u := &User{}
	err := db.Pool.QueryRow(ctx, `
		SELECT 
			id, phone, password, refresh_token, refresh_token_at, role, status, name, email, avatar, bio,
			social_media_1_name, social_media_1_url, social_media_2_name, social_media_2_url,
			rating_avg, rating_count, verified, created_at
		FROM users
		WHERE refresh_token=$1
	`, token).Scan(
		&u.ID, &u.Phone, &u.Password, &u.RefreshToken, &u.RefreshTokenAt,
		&u.Role, &u.Status, &u.Name, &u.Email, &u.Avatar, &u.Bio,
		&u.SocialMedia1Name, &u.SocialMedia1URL, &u.SocialMedia2Name, &u.SocialMedia2URL,
		&u.RatingAvg, &u.RatingCount, &u.Verified, &u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Create a new user
func CreateUser(ctx context.Context, u *User) error {
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO users (
			phone, password, name, email, avatar, bio, social_media_1_name, social_media_1_url,
			social_media_2_name, social_media_2_url
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`, u.Phone, u.Password, u.Name, u.Email, u.Avatar, u.Bio, u.SocialMedia1Name, u.SocialMedia1URL,
		u.SocialMedia2Name, u.SocialMedia2URL)
	return err
}

// Update refresh token
func UpdateUserRefreshToken(ctx context.Context, userID int64, token string) error {
	var err error
	if token == "" {
		_, err = db.Pool.Exec(ctx, `
			UPDATE users
			SET refresh_token = '', refresh_token_at = NULL
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

// Update user info
func UpdateUser(ctx context.Context, u *User) error {
	_, err := db.Pool.Exec(ctx, `
		UPDATE users
		SET phone=$1, name=$2, email=$3, avatar=$4, bio=$5,
		    social_media_1_name=$6, social_media_1_url=$7,
		    social_media_2_name=$8, social_media_2_url=$9,
		    verified=$10, status=$11
		WHERE id=$12
	`, u.Phone, u.Name, u.Email, u.Avatar, u.Bio,
		u.SocialMedia1Name, u.SocialMedia1URL,
		u.SocialMedia2Name, u.SocialMedia2URL,
		u.Verified, u.Status, u.ID)
	return err
}

package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `env:"HTTP_ADDRESS"`
}

type Config struct {
	APP_ENV string `env:"APP_ENV"`
	DB_URL  string `env:"DB_URL"`
	HTTPServer
	JWTKey             string `env:"JWT_KEY"`
	AccessTokenTTL     int    `env:"ACCESS_TOKEN_TTL_MIN" env-default:"15"`
	RefreshTokenTTL    int    `env:"REFRESH_TOKEN_TTL_DAYS" env-default:"30"`
	SuperAdminEmail    string `env:"SUPERADMIN_EMAIL"`
	SuperAdminPassword string `env:"SUPERADMIN_PASSWORD"`
}

func LoadConfig() *Config {
	var cfg Config

	var envPath string
	flag.StringVar(&envPath, "config", "", "path to .env file")
	flag.Parse()

	if envPath == "" {
		envPath = os.Getenv("CONFIG_PATH")
	}

	if envPath == "" {
		envPath = "config/dev.env"
	}

	if err := cleanenv.ReadConfig(envPath, &cfg); err != nil {
		log.Fatalf("Cannot read config from %s: %v", envPath, err)
	}

	if cfg.JWTKey == "" || cfg.SuperAdminEmail == "" || cfg.SuperAdminPassword == "" || cfg.DB_URL == "" {
		log.Fatal("JWT_KEY, SUPERADMIN_EMAIL, SUPERADMIN_PASSWORD, and DB_URL must be set")
	}

	return &cfg
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/models"
	"backend/internal/routes"
	"backend/internal/utils"
)

func main() {
	cfg := config.LoadConfig()
	log.Printf("Starting Bhinno backend in %s mode...", cfg.APP_ENV)

	db := db.Init(cfg)
	defer db.Close()

	superadminPasswordHash := utils.HashPassword(cfg.SuperAdminPassword)
	models.EnsureSuperAdmin(cfg.SuperAdminEmail, superadminPasswordHash)

	utils.InitJWT(cfg.JWTKey, cfg.AccessTokenTTL)
	utils.InitRefreshTokenTTL(cfg.RefreshTokenTTL)

	mux := routes.RegisterRoutes()

	srv := &http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: mux,
	}

	go func() {
		log.Printf("HTTP server running at %s", cfg.HTTPServer.Address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	gracefulShutdown(srv, 10*time.Second)
}

func gracefulShutdown(srv *http.Server, timeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quit
	log.Printf("Received signal %s, shutting down...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}

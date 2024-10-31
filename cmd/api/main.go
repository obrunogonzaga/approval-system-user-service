package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/obrunogonzaga/go-template/configs"
	"github.com/obrunogonzaga/go-template/internal/handler"
	"github.com/obrunogonzaga/go-template/internal/repository"
	"github.com/obrunogonzaga/go-template/internal/usecase"
	"github.com/obrunogonzaga/go-template/pkg/logger"
)

func main() {
	// Initialize logger
	zapLogger := logger.NewLogger()
	defer zapLogger.Sync()

	// Load configuration
	config := loadConfig()

	// Initialize database
	db, err := initializeDB(config.Database)
	if err != nil {
		zapLogger.Fatal("Failed to initialize database", zap.Error(err))
	}
	defer db.Close()

	// Initialize router
	router := gin.Default()

	// Initialize dependencies
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	// Setup API routes
	api := router.Group("/api/v1")
	{
		userHandler.RegisterRoutes(api)
	}

	// Setup server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.App.Port),
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		zapLogger.Info("Starting server", zap.String("port", config.App.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zapLogger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zapLogger.Info("Shutting down server...")

	// Give outstanding requests a deadline for completion
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zapLogger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	zapLogger.Info("Server exiting")
}

func loadConfig() *configs.Config {
	// Here you would implement your configuration loading logic
	// This is a simple example - you might want to use environment variables or files
	return &configs.Config{
		App: configs.AppConfig{
			Port: getEnvOrDefault("APP_PORT", "8080"),
			Env:  getEnvOrDefault("APP_ENV", "development"),
		},
		Database: configs.DatabaseConfig{
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvOrDefault("DB_PORT", "5432"),
			User:     getEnvOrDefault("DB_USER", "postgres"),
			Password: getEnvOrDefault("DB_PASSWORD", "postgres"),
			DBName:   getEnvOrDefault("DB_NAME", "myapp"),
		},
	}
}

func initializeDB(config configs.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test database connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return db, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

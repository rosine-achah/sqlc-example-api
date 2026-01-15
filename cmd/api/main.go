package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/Iknite-Space/sqlc-example-api/api"
	"github.com/Iknite-Space/sqlc-example-api/db/repo"
	"github.com/ardanlabs/conf/v3"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBUser      string `conf:"env:DB_USER,required"`
	DBPassword  string `conf:"env:DB_PASSWORD,required,mask"`
	DBHost      string `conf:"env:DB_HOST,required"`
	DBPort      uint16 `conf:"env:DB_PORT,required"`
	DBName      string `conf:"env:DB_NAME,required"`
	TLSDisabled bool   `conf:"env:DB_TLS_DISABLED"`
}

type Config struct {
	ListenPort     uint16 `conf:"env:LISTEN_PORT,required"`
	MigrationsPath string `conf:"env:MIGRATIONS_PATH,required"`
	DB             DBConfig
}

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	config := Config{}

	if err := loadConfig(&config); err != nil {
		return err
	}

	dbURL := getPostgresConnectionURL(config.DB)
	db, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}
	defer db.Close()

	if err := repo.Migrate(dbURL, config.MigrationsPath); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	querier := repo.New(db)
	router := gin.Default()

	// Initialize order handler (CampayClient inside)
	orderHandler := api.NewOrderHandler(querier)
	orderHandler.WireRoutes(router)

	msgHandler := api.NewMessageHandler(querier)
	msgHandler.WireHttpHandler(router)

	fmt.Printf("Server listening on port %d\n", config.ListenPort)
	return router.Run(fmt.Sprintf(":%d", config.ListenPort))
}

func loadConfig(cfg *Config) error {
	if _, err := os.Stat(".env"); err != nil {
		return fmt.Errorf("error loading env: %v", err)
	}
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("failed to load .env: %w", err)
	}
	if _, err := conf.Parse("", cfg); err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			return err
		}
		return err
	}
	fmt.Printf("Loaded config: %+v\n", cfg)
	return nil
}

func getPostgresConnectionURL(cfg DBConfig) string {
	query := url.Values{}
	if cfg.TLSDisabled {
		query.Add("sslmode", "disable")
	} else {
		query.Add("sslmode", "require")
	}

	return (&url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.DBUser, cfg.DBPassword),
		Host:     fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort),
		Path:     cfg.DBName,
		RawQuery: query.Encode(),
	}).String()
}

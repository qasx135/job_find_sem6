package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
	"time"
)

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func NewPostgresConfig(username, password, host, port, database string) *PostgresConfig {
	return &PostgresConfig{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

func NewClient(ctx context.Context, cfg *PostgresConfig, maxAttempts int, delay time.Duration) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	log.Printf("Connecting to DB: %s", dsn)
	var pool *pgxpool.Pool
	err := DoWithTries(func() error {
		ctxTimeout, cancel := context.WithTimeout(ctx, delay)
		defer cancel()
		config, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			return err
		}
		pool, err = pgxpool.NewWithConfig(ctxTimeout, config)
		if err != nil {
			log.Println("Waiting for DB to become available...")
			return err
		}
		var n int
		if err := pool.QueryRow(context.Background(), "SELECT 1").Scan(&n); err != nil {
			log.Println("DB not ready yet:", err)
			return err
		}
		return nil
	}, maxAttempts, delay)
	if err != nil {
		log.Fatalf("Could not connect to DB after several attempts: %v", err)
		return nil, err
	}
	log.Println("Connected to DB and DB is ready!")
	time.Sleep(5 * time.Second)
	if err := Migrate(cfg); err != nil {
		log.Fatalf("Migration failed: %v", err)
		return nil, err
	}
	return pool, nil
}

func DoWithTries(fn func() error, maxAttempts int, delay time.Duration) error {
	var err error
	for maxAttempts > 0 {
		if err = fn(); err == nil {
			return nil
		}
		time.Sleep(delay)
		maxAttempts--
	}
	return err
}

func Migrate(cfg *PostgresConfig) error {
	goose.SetVerbose(true)
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username,
		cfg.Password, cfg.Host, cfg.Port,
		cfg.Database)
	sql, err := goose.OpenDBWithDriver("postgres", dbUrl)
	if err != nil {
		log.Fatalf("goose open: %w", err)
		return err
	}
	defer sql.Close()
	if err := goose.Up(sql, "./db/migrations"); err != nil {
		log.Fatalf("goose up: %v", err)
		return err
	}
	return nil
}

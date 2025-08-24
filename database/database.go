package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDatabase() *pgxpool.Pool {
	dsn := os.Getenv("DATABASE_URL")
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Failed to parse DB config: %v", err)
	}

	// kill the stmt cache if you want to avoid the collision entirely
	// cfg.ConnConfig.StatementCacheCapacity = 0
	// cfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	ctx := context.Background()

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	var version string
	if err := pool.QueryRow(ctx, "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	log.Println("Connected to:", version)

	return pool
}

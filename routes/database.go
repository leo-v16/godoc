package routes

import "github.com/jackc/pgx/v5/pgxpool"

type DB struct {
	POOL *pgxpool.Pool
}

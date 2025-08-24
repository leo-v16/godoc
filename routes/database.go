package routes

import "github.com/jackc/pgx/v5"

type DB struct {
	CONN *pgx.Conn
}

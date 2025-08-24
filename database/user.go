package database

import (
	"context"
	"errors"
	"fmt"
	"godoc/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateUser(conn *pgx.Conn, user *models.User) (int, error) {
	var user_id int
	err := conn.QueryRow(context.Background(),
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id;",
		user.Username, user.Password).Scan(&user_id)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				println("USERNAME ALEADY EXISTS")
				return 0, fmt.Errorf("duplicate")
			}
		}
		println("USER CREATION FAIELD| ERROR: ", err.Error())
		return 0, err
	} else {
		println("USER CREATED| ID: ", user_id)
		return user_id, nil
	}
}

func LogInUser(conn *pgx.Conn, user *models.User) (int, error) {
	var password string
	var user_id int

	err := conn.QueryRow(context.Background(),
		"SELECT password, id FROM users WHERE username = $1",
		user.Username).Scan(&password, &user_id)

	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, fmt.Errorf("not found")
		}
		return 0, err
	}

	if password == user.Password {
		println("Log In Success")
		return user_id, nil
	} else {
		return 0, fmt.Errorf("wrong password")
	}
}

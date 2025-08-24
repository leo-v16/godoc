package database

import (
	"context"
	"godoc/models"

	"github.com/jackc/pgx/v5"
)

func CreateQuery(conn *pgx.Conn, query *models.Query) (int, error) {
	var query_id int
	err := conn.QueryRow(context.Background(),
		"INSERT INTO query (user_id, text) VALUES ($1, $2) RETURNING id;",
		query.UserId, query.Text).Scan(&query_id)
	if err != nil {
		println("QUERY NOT CREATED | ERROR: ", err.Error())
		return 0, err
	} else {
		println("QUERY CREATED | ID: ", query_id)
		return query_id, nil
	}
}

func GetAllQuery(conn *pgx.Conn) ([]models.Query, error) {
	var queryArray []models.Query
	rows, err := conn.Query(context.Background(), "SELECT id, user_id, text FROM query;")
	if err != nil {
		println("COULD NOT LOAD QUERY | ERROR: ", err.Error())
		return queryArray, err
	}
	for rows.Next() {
		var query models.Query
		rows.Scan(&query.Id, &query.UserId, &query.Text)
		queryArray = append(queryArray, query)
		println("ID:", query.Id, "|USER_ID:", query.UserId, "|TEXT:", query.Text)
	}
	return queryArray, nil
}

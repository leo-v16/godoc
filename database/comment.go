package database

import (
	"context"
	"godoc/models"

	"github.com/jackc/pgx/v5"
)

func CreateComment(conn *pgx.Conn, comment *models.Comment) (int, error) {
	var comment_id int
	if err := conn.QueryRow(context.Background(),
		"INSERT INTO comment (query_id, user_id, text) VALUES ($1, $2, $3) RETURNING id;",
		comment.QueryId, comment.UserId, comment.Text).Scan(&comment_id); err == nil {
		println("COMMENT CREATION SUCCESS")
		return comment_id, nil
	} else {
		println("COMMENT NOT CREATED | ERROR: ", err.Error())
		return 0, err
	}
}

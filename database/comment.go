package database

import (
	"context"
	"godoc/models"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateComment(ctx context.Context, pool *pgxpool.Pool, comment *models.Comment) (int, error) {
	var commentID int
	err := pool.QueryRow(ctx,
		`INSERT INTO comment (query_id, user_id, text) 
         VALUES ($1, $2, $3) RETURNING id;`,
		comment.QueryId, comment.UserId, comment.Text,
	).Scan(&commentID)

	if err != nil {
		log.Printf("COMMENT NOT CREATED | ERROR: %v", err)
		return 0, err
	}

	log.Printf("COMMENT CREATION SUCCESS | id=%d", commentID)
	return commentID, nil
}

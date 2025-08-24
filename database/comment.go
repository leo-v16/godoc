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

func GetAllCommentOfQuery(ctx context.Context, pool *pgxpool.Pool, queryId int) ([]models.Comment, error) {
	var commentArray []models.Comment
	rows, err := pool.Query(ctx,
		`SELECT id, query_id, user_id, text FROM comment WHERE query_id = $1;`,
		queryId)
	if err != nil {
		println("COULD NOT LOAD COMMENT | ERROR: ", err.Error())
		return commentArray, err
	}
	for rows.Next() {
		var comment models.Comment
		rows.Scan(&comment.Id, &comment.QueryId, &comment.UserId, &comment.Text)
		commentArray = append(commentArray, comment)
	}
	return commentArray, nil
}

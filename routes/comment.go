package routes

import (
	"godoc/database"
	"godoc/models"

	"github.com/gin-gonic/gin"
	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRouteComment(router *gin.RouterGroup, pool *pgxpool.Pool) {
	db := &DB{POOL: pool}
	router.POST("/create", db.CreateCommentEndPoint)
}

func (D *DB) CreateCommentEndPoint(ctx *gin.Context) {
	var comment models.Comment
	if err := ctx.ShouldBind(&comment); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
	}

	comment_id, err := database.CreateComment(D.POOL, &comment)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Couldn't create comment",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"id":      comment_id,
		"message": "Comment created succesfully",
	})
}

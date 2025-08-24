package routes

import (
	"godoc/database"
	"godoc/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRouteComment(router *gin.RouterGroup, pool *pgxpool.Pool) {
	db := &DB{POOL: pool}
	router.POST("/create", db.CreateCommentEndPoint)
	router.POST("/getallofquery")
}

func (D *DB) CreateCommentEndPoint(ctx *gin.Context) {
	var comment models.Comment

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	commentID, err := database.CreateComment(ctx.Request.Context(), D.POOL, &comment)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Couldn't create comment",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"id":      commentID,
		"message": "Comment created successfully",
	})
}

func (D *DB) GetAllCommentOfQuery(ctx *gin.Context) {
	var query models.Query

	if err := ctx.ShouldBindJSON(&query); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	commentArray, err := database.GetAllCommentOfQuery(ctx.Request.Context(), D.POOL, query.Id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Couldn't fetch query",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"count":        len(commentArray),
		"comment_list": commentArray,
	})
}

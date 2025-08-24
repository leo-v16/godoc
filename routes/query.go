package routes

import (
	"godoc/database"
	"godoc/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRouteQuery(router *gin.RouterGroup, pool *pgxpool.Pool) {
	db := &DB{POOL: pool}
	router.POST("/create", db.CreateQueryEndPoint)
	router.GET("/getall", db.GetAllQueryEndPoint)
}

func (D *DB) CreateQueryEndPoint(ctx *gin.Context) {
	var query models.Query
	if err := ctx.ShouldBind(&query); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	query_id, err := database.CreateQuery(D.POOL, &query)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Query creation failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"id":      query_id,
		"message": "Query created succesfully",
	})
}

func (D *DB) GetAllQueryEndPoint(ctx *gin.Context) {
	queryArray, err := database.GetAllQuery(D.POOL)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Couldn't fetch query",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"count":      len(queryArray),
		"query_list": queryArray,
	})
}

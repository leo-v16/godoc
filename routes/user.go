package routes

import (
	"godoc/database"
	"godoc/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRouteUser(router *gin.RouterGroup, pool *pgxpool.Pool) {
	db := &DB{POOL: pool}
	router.POST("/signin", db.CreateUserEndPoint)
	router.POST("/login", db.LogInUserEndPoint)
}

func (D *DB) CreateUserEndPoint(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		println("PAYLOAD BINDING ERROR : ", err.Error())
		ctx.JSON(400, gin.H{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	user_id, err := database.CreateUser(ctx.Request.Context(), D.POOL, &user)
	if err != nil {
		if err.Error() == "duplicate" {
			ctx.JSON(409, gin.H{
				"message": "Username already exists",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "User creation failed",
				"error":   err.Error(),
			})
		}
		return
	}

	user.Id = user_id
	ctx.JSON(201, gin.H{
		"id":      user.Id,
		"message": "User created succesfully",
	})
}

func (D *DB) LogInUserEndPoint(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	user_id, err := database.LogInUser(ctx.Request.Context(), D.POOL, &user)
	if err != nil {
		if err.Error() == "not found" {
			ctx.JSON(404, gin.H{
				"message": "Username not found",
			})
		} else if err.Error() == "wrong password" {
			ctx.JSON(401, gin.H{
				"message": "Incorrect password",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Login failed",
				"error":   err.Error(),
			})
		}
		return
	}

	ctx.JSON(200, gin.H{
		"id":      user_id,
		"message": "Login Success",
	})
}

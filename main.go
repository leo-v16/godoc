package main

import (
	"context"
	"godoc/database"
	"godoc/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "github.com/gin-gonic/gin"
)

type PayLoad struct {
	Data string `json:"data"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	conn := database.ConnectDatabase()
	defer conn.Close(context.Background())

	router := gin.Default()
	user_route := router.Group("/user")
	routes.RegisterRouteUser(user_route, conn)
	routes.RegisterRouteQuery(router.Group("/query"), conn)
	routes.RegisterRouteComment(router.Group("/comment"), conn)
	router.Run(":8080")

	// user := &models.User{Username: "**kjhsdf", Password: "GG"}
	// database.LogInUser(conn, user)

	// for _, q := range database.GetAllQuery(conn) {
	// 	println("ID:", q.Id, "|USER_ID:", q.UserId, "|TEXT:", q.Text)
	// }

	// database.CreateComment(conn, &models.Comment{UserId: 2, QueryId: 1, Text: "This is le"})

	// if row := conn.QueryRow(context.Background(), "SELECT * from users;"); true {
	// 	var user User
	// 	err := row.Scan(&user.Id, &user.Username, &user.Password)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	println("ID: ", user.Id, "|USERNAME: ", user.Username, "|PASSWORD: ", user.Password)
	// }

	// rows, err := conn.Query(context.Background(), "SELECT id, username, password FROM users;")
	// if err != nil {
	// 	panic(err.Error()) // or handle error gracefully
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var user User
	// 	err := rows.Scan(&user.Id, &user.Username, &user.Password)
	// 	if err != nil {
	// 		panic(err) // or handle error gracefully
	// 	}
	// 	println("ID: ", user.Id, "|Username: ", user.Username, "|Password: ", user.Password)
	// }

	// // Check for errors from iterating over rows
	// if err = rows.Err(); err != nil {
	// 	panic(err) // or handle error gracefully
	// }

	// router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// router.POST("/ping", func(ctx *gin.Context) {
	// 	var data PayLoad
	// 	ctx.ShouldBind(&data)
	// 	ctx.JSON(200, gin.H{
	// 		"message": "Hello " + data.Data,
	// 	})
	// })

	// router.Run()
}

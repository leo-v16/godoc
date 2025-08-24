package main

import (
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
	defer conn.Close()

	router := gin.Default()
	user_route := router.Group("/user")
	routes.RegisterRouteUser(user_route, conn)
	routes.RegisterRouteQuery(router.Group("/query"), conn)
	routes.RegisterRouteComment(router.Group("/comment"), conn)
	router.Run(":8080")
}

package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type Query struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

type Comment struct {
	Id      int    `json:"id"`
	QueryId int    `json:"query_id" binding:"required"`
	UserId  int    `json:"user_id" binding:"required"`
	Text    string `json:"text" binding:"required"`
}

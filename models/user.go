package models

type User struct {
	Id   int
	Name string
}

type UserInput struct {
	Name string `json:"name" binding:"required"`
}

package models

type Lesson struct {
	Id     int
	UserId int 
	Name   string
}

type LessonInput struct {
	UserId int    `json:"user_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

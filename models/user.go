package models

type User struct {
	Id   int    `gorm:"primaryKey;column:id"`
	Name string `gorm:"column:name"`
}

type UserInput struct {
	Name string `json:"name" binding:"required"`
}

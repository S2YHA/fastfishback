package models

type Lesson struct {
	Id     int    `gorm:"primaryKey;column:id"`
	UserId int    `gorm:"column:user_id" json:"user_id"`
	Name   string `gorm:"column:name"`
	Words  []Word `gorm:"foreignKey:LessonId"`
}

// type LessonFull struct {
// 	Id     int
// 	UserId int    `json:"user_id" binding:"required"`
// 	Name   string `json:"name"`
// 	Words  []Word `json:"words"`
// }

// type LessonInput struct {
// 	UserId int         `json:"user_id" binding:"required"`
// 	Name   string      `json:"name" binding:"required"`
// 	Words  []WordInput `json:"words"`
// }

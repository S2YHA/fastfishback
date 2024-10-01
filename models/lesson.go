package models

type Lesson struct {
	Id     int    `gorm:"primaryKey;column:id"`
	UserId int    `gorm:"column:user_id" json:"user_id"`
	Name   string `gorm:"column:name"`
	Words  []Word `gorm:"foreignKey:LessonId"`
}

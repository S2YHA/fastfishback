package models

type Word struct {
	Id         int    `gorm:"primaryKey;column:id"`
	LessonId   int    `gorm:"column:lesson_id"`
	FirstWord  string `gorm:"column:first_word" json:"first_word"`
	SecondWord string `gorm:"column:second_word" json:"second_word"`
}

// type WordInput struct {
// 	LessonId   int    `gorm:"foreignKey:LessonId"`
// 	FirstWord  string `json:"first_word"`
// 	SecondWord string `json:"second_word"`
// }

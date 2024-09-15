package services

import (
	"fastfishback/models"
	"fastfishback/repositories"
)

func GetAllLessons() ([]models.Lesson, error) {
	return repositories.FindAllLessons()
}

func GetLessonByID(id uint) (*models.Lesson, error) {
	return repositories.FindLessonByID(id)
}

func CreateLesson(LessonInput *models.LessonInput) (*models.Lesson, error) {
	return repositories.CreateLesson(LessonInput)
}

func UpdateLesson(id uint, LessonInput models.LessonInput) (*models.Lesson, error) {
	return repositories.UpdateLesson(id, LessonInput)
}

func DeleteLesson(id uint) error {
	return repositories.DeleteLesson(id)
}

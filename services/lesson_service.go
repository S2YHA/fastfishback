package services

import (
	"fastfishback/config"
	"fastfishback/models"
	"fastfishback/repositories"
	"fmt"

	"gorm.io/gorm"
)

func GetAllLessons() ([]models.Lesson, error) {
	lessons, err := repositories.FindAllLessons()
	if err != nil {
		return nil, err
	}

	for id := range lessons {
		words, err := repositories.FindWordsByLessonID(uint(lessons[id].Id))
		if err != nil {
			return nil, err
		}

		lessons[id].Words = words
	}

	return lessons, nil
}

func GetLessonByID(id uint) (*models.Lesson, error) {
	lesson, err := repositories.FindLessonByID(id)
	if err != nil {
		return nil, err
	}

	words, err := repositories.FindWordsByLessonID(id)
	if err != nil {
		return nil, err
	}

	lesson.Words = words

	return lesson, err
}

func CreateLesson(lesson *models.Lesson) (*models.Lesson, error) {
	var newLesson *models.Lesson

	fmt.Println(lesson)

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		newLesson, err := repositories.CreateLesson(lesson)
		if err != nil {
			return err
		}

		for i := range lesson.Words {
			lesson.Words[i].LessonId = newLesson.Id
		}

		if err := repositories.CreateWords(lesson.Words); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return newLesson, nil
}

func UpdateLesson(id uint, LessonInput models.Lesson) (*models.Lesson, error) {
	return repositories.UpdateLesson(id, LessonInput)
}

func DeleteLesson(id uint) error {
	return repositories.DeleteLesson(id)
}

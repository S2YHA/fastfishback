package services

import (
	"fastfishback/config"
	"fastfishback/models"
	"fastfishback/repositories"

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

func UpdateLesson(id uint, lessonInput models.Lesson) (*models.Lesson, error) {
	var existingLesson *models.Lesson

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if id != 0 {
			if _, err := repositories.FindLessonByID(id); err != nil {
				return err
			}

			if _, err := repositories.UpdateLesson(id, lessonInput); err != nil {
				return err
			}
		}

		if len(lessonInput.Words) != 0 {
			if err := repositories.DeleteWordsByLessonId(id); err != nil {
				return err
			}

			for i := range lessonInput.Words {
				lessonInput.Words[i].LessonId = int(id)
			}

			if err := repositories.CreateWords(lessonInput.Words); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return existingLesson, nil
}

func DeleteLesson(id uint) error {
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := repositories.DeleteWordsByLessonId(id); err != nil {
			return err
		}

		if err := repositories.DeleteLesson(id); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

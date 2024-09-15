package repositories

import (
	"fastfishback/config"
	"fastfishback/models"
)

func FindAllLessons() ([]models.Lesson, error) {
	var lessons []models.Lesson
	if err := config.DB.Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
}

func FindLessonByID(id uint) (*models.Lesson, error) {
	var lessons models.Lesson
	if err := config.DB.First(&lessons, id).Error; err != nil {
		return nil, err
	}
	return &lessons, nil
}

func CreateLesson(lessonInput *models.LessonInput) (*models.Lesson, error) {
	lesson := models.Lesson{
		UserId: lessonInput.UserId,
		Name:   lessonInput.Name,
	}
	if err := config.DB.Create(&lesson).Error; err != nil {
		return nil, err
	}
	return &lesson, nil
}

func UpdateLesson(id uint, lessonInput models.LessonInput) (*models.Lesson, error) {
	lesson := models.Lesson{
		UserId: lessonInput.UserId,
		Name:   lessonInput.Name,
	}
	if err := config.DB.Model(&models.Lesson{}).Where("id = ?", id).Updates(&lesson).Error; err != nil {
		return nil, err
	}
	return &lesson, nil
}

func DeleteLesson(id uint) error {
	if err := config.DB.Where("id = ?", id).Delete(&models.Lesson{}).Error; err != nil {
		return err
	}
	return nil
}

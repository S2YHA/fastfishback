package repositories

import (
	"fastfishback/config"
	"fastfishback/models"
)

func CreateWords(words []models.Word) error {
	return config.DB.Create(&words).Error
}

func FindWordsByLessonID(id uint) ([]models.Word, error) {
	var words []models.Word
	if err := config.DB.Where("lesson_id = ?", id).Find(&words).Error; err != nil {
		return nil, err
	}
	return words, nil
}

func DeleteWordsByLessonId(lessonId uint) error {
	return config.DB.Where("lesson_id = ?", lessonId).Delete(&models.Word{}).Error
}

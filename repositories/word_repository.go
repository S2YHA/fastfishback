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

// func UpdateLesson(id uint, lessonInput models.LessonInput) (*models.Lesson, error) {
// 	lesson := models.Lesson{
// 		UserId: lessonInput.UserId,
// 		Name:   lessonInput.Name,
// 	}
// 	if err := config.DB.Model(&models.Lesson{}).Where("id = ?", id).Updates(&lesson).Error; err != nil {
// 		return nil, err
// 	}
// 	return &lesson, nil
// }

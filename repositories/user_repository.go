package repositories

import (
	"fastfishback/config"
	"fastfishback/models"
)

func FindAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func FindUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(userInput *models.UserInput) (*models.User, error) {
	user := models.User{
		Name: userInput.Name,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id uint, userInput models.UserInput) (*models.User, error) {
	user := models.User{
		Name: userInput.Name,
	}
	if err := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(id uint) error {
	if err := config.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

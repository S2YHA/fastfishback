package services

import (
	"fastfishback/models"
	"fastfishback/repositories"
)

func GetAllUsers() ([]models.User, error) {
	return repositories.FindAllUsers()
}

func GetUserByID(id uint) (*models.User, error) {
	return repositories.FindUserByID(id)
}

func CreateUser(userInput *models.UserInput) (*models.User, error) {
	return repositories.CreateUser(userInput)
}

func UpdateUser(id uint, userInput models.UserInput) (*models.User, error) {
	return repositories.UpdateUser(id, userInput)
}

func DeleteUser(id uint) error {
	return repositories.DeleteUser(id)
}

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

// Other business logic (create user, update user, etc.)

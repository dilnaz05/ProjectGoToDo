package services

import (
	"errors"
	"todo-app/backend/db"
	"todo-app/backend/models"
	"todo-app/backend/utils"
)

func Login(username, password string) (string, error) {
	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	if err := utils.CheckPasswordHash(user.Password, password); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func Register(username, password string) error {
	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err == nil {
		return errors.New("username already taken")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user = models.User{Username: username, Password: hashedPassword}
	if err := db.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

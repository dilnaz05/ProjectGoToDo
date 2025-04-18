package services

import (
	"todo-app/backend/db"
	"todo-app/backend/models"
)

func GetTodosByUserID(userID uint) ([]models.Todo, error) {
	var todos []models.Todo
	if err := db.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func CreateTodo(userID uint, message string) (models.Todo, error) {
	todo := models.Todo{
		UserID:   userID,
		Message:  message,
		Complete: false,
	}

	if err := db.DB.Create(&todo).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

func GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	if err := db.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func UpdateTodo(userID uint, id uint, message string, complete bool) (models.Todo, error) {
	var todo models.Todo
	if err := db.DB.Where("user_id = ? AND id = ?", userID, id).First(&todo).Error; err != nil {
		return todo, err
	}

	todo.Message = message
	todo.Complete = complete

	if err := db.DB.Save(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func DeleteTodo(userID uint, id uint) error {
	if err := db.DB.Where("user_id = ? AND id = ?", userID, id).Delete(&models.Todo{}).Error; err != nil {
		return err
	}
	return nil
}

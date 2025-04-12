package services

import (
	"todo-app-backend/internal/db"
	"todo-app-backend/internal/models"
)

func CreateTodo(message string) (*models.Todo, error) {
	todo := &models.Todo{Message: message, Complete: false}
	if err := db.DB.Create(todo).Error; err != nil {
		return nil, err
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

func UpdateTodo(id uint, message string, complete bool) (*models.Todo, error) {
	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}

	todo.Message = message
	todo.Complete = complete
	if err := db.DB.Save(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func DeleteTodo(id uint) error {
	if err := db.DB.Delete(&models.Todo{}, id).Error; err != nil {
		return err
	}
	return nil
}

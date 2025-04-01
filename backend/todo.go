package main

import (
	"errors"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Message  string `json:"message"`
	Complete bool   `json:"complete"`
}

func Get() ([]Todo, error) {
	var todos []Todo
	result := DB.Find(&todos)
	return todos, result.Error
}

func GetItem(id uint) (*Todo, error) {
	var todo Todo
	result := DB.First(&todo, id)
	if result.Error != nil {
		return nil, errors.New("not found")
	}
	return &todo, nil
}

func Add(message string) (*Todo, error) {
	todo := Todo{Message: message, Complete: false}
	result := DB.Create(&todo)
	return &todo, result.Error
}

func Delete(id uint) error {
	result := DB.Delete(&Todo{}, id)
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return result.Error
}

func Complete(id uint) error {
	var todo Todo
	result := DB.First(&todo, id)
	if result.Error != nil {
		return errors.New("not found")
	}
	todo.Complete = true
	DB.Save(&todo)
	return nil
}

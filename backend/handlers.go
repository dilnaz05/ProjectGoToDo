package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTodoListHandler(c *gin.Context) {
	todos, err := Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func AddTodoHandler(c *gin.Context) {
	var request struct {
		Message string `json:"message"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	todo, err := Add(request.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add todo"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": todo.ID})
}

func DeleteTodoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func CompleteTodoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := Complete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.Status(http.StatusOK)
}

func GetTodoItemHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := GetItem(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodoItemHandler(c *gin.context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var request struct {
		Message  string `json:"message"`
		Complete bool   `json:"complete"`
	}

	if err := c.ShouldBindJSON(&request); err != nill {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid input"})
		return
	}

	updatedTodo, err := Update(uint(id), request.Message, request.Complete)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

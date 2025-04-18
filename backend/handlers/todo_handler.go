package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-app/backend/services"
)

func CreateTodoHandler(c *gin.Context) {
	// Авторизацияланған пайдаланушының userID алу
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Сұраныстан Todo мәліметтерін алу
	var req struct {
		Message string `json:"message"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Пайдаланушыға тиісті Todo жасау
	todo, err := services.CreateTodo(userID.(uint), req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func GetTodosHandler(c *gin.Context) {
	// Авторизацияланған пайдаланушының userID алу
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Пайдаланушының Todo-ларын алу
	todos, err := services.GetTodosByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func UpdateTodoHandler(c *gin.Context) {
	// Авторизацияланған пайдаланушының userID алу
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// URL параметрінен Todo ID алу
	id, _ := strconv.Atoi(c.Param("id"))

	// Сұраныстан жаңа мәліметтерді алу
	var req struct {
		Message  string `json:"message"`
		Complete bool   `json:"complete"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Пайдаланушыға тиісті Todo жаңарту
	todo, err := services.UpdateTodo(userID.(uint), uint(id), req.Message, req.Complete)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodoHandler(c *gin.Context) {
	// Авторизацияланған пайдаланушының userID алу
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// URL параметрінен Todo ID алу
	id, _ := strconv.Atoi(c.Param("id"))

	// Пайдаланушыға тиісті Todo жою
	err := services.DeleteTodo(userID.(uint), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.Status(http.StatusNoContent)
}

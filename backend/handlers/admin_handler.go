package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/backend/db"
	"todo-app/backend/models"
)

func AdminDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the admin dashboard"})
}

func GetAllUsers(c *gin.Context) {
	currentUserID := c.GetUint("userID")

	var users []models.User
	if err := db.DB.Where("id != ?", currentUserID).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	db.DB.Preload("User").Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func DeleteUser(c *gin.Context) {
	currentUserID := c.GetUint("userID")
	targetID := c.Param("id")

	if fmt.Sprintf("%d", currentUserID) == targetID {
		c.JSON(http.StatusForbidden, gin.H{"error": "you cannot delete yourself"})
		return
	}

	db.DB.Delete(&models.User{}, targetID)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTodoListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Get())
}

func AddTodoHandler(c *gin.Context) {
	var request struct {
		Message string `json:"message"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	id := Add(request.Message)
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func DeleteTodoHandler(c *gin.Context) {
	id := c.Param("id")
	if err := Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func CompleteTodoHandler(c *gin.Context) {
	var request struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := Complete(request.ID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.Status(http.StatusOK)
}

func GetTodoItemHandler(c *gin.Context) {
	id := c.Param("id")
	todo, err := GetItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

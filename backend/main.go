package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/todo", GetTodoListHandler)
	r.POST("/todo", AddTodoHandler)
	r.DELETE("/todo/:id", DeleteTodoHandler)
	r.PUT("/todo", CompleteTodoHandler)
	r.GET("/todo/:id", GetTodoItemHandler)

	r.Run(":3000")
}

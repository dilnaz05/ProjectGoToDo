package main

import (
	"github.com/gin-gonic/gin"
	"todo-app-backend/internal/db"
	"todo-app-backend/internal/handlers"
	"todo-app-backend/internal/middleware"
)

func main() {
	db.ConnectDB()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)

		protected := api.Group("/")
		protected.Use(middleware.JWTAuth())
		{
			protected.GET("/me", handlers.Me)

			protected.GET("/todos", handlers.GetTodos)
			protected.POST("/todos", handlers.CreateTodo)
			protected.PUT("/todos/:id", handlers.UpdateTodo)
			protected.DELETE("/todos/:id", handlers.DeleteTodo)
		}
	}

	r.Run(":8080")
}

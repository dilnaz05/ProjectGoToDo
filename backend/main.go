package main

import (
	"github.com/gin-gonic/gin"
	"todo-app/backend/db"
	"todo-app/backend/handlers"
	"todo-app/backend/middleware"
)

func main() {
	db.InitDB() // ✅

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", handlers.RegisterHandler) // ✅
		api.POST("/login", handlers.LoginHandler)       // ✅

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware()) // ✅

		{
			protected.GET("/me", handlers.MeHandler)                   // ✅
			protected.GET("/todos", handlers.GetTodosHandler)          // ✅
			protected.POST("/todos", handlers.CreateTodoHandler)       // ✅
			protected.PUT("/todos/:id", handlers.UpdateTodoHandler)    // ✅
			protected.DELETE("/todos/:id", handlers.DeleteTodoHandler) // ✅
		}
	}

	r.Run(":8080")
}

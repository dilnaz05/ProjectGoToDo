package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"todo-app-backend/internal/services"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=Butterfly dbname=tododb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	fmt.Println("Database connection successful!")
	DB = db

	//migration
	err = db.AutoMigrate(&services.main{})
	if err != nil {
		log.Fatal("Migration error: ", err)
	}
}

package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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

	err = db.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatal("Migration error: ", err)
	}
}

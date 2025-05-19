package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	//"os/exec"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	fmt.Println("Database connection successful!")
	DB = db

	//runSQLMigration()
	//err = db.AutoMigrate(&models.User{}, &models.Todo{})
}

//func runSQLMigration() {
//	cmd := exec.Command(
//		"C:\\Program Files\\PostgreSQL\\16\\bin\\psql.exe",
//		"-U", "postgres",
//		"-d", "tododb",
//		"-f", "C:\\Users\\Lenovo\\GolandProjects\\todo-app\\backend\\db\\migration\\init_schema.sql",
//	)
//
//	output, err := cmd.CombinedOutput()
//	if err != nil {
//		log.Fatalf("SQL migration error: %s\n%s", err, output)
//	}
//	fmt.Println("SQL migration applied successfully")
//}

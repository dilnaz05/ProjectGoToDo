package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os/exec"
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

	runSQLMigration()
	//err = db.AutoMigrate(&models.User{}, &models.Todo{})
}

func runSQLMigration() {
	cmd := exec.Command(
		"C:\\Program Files\\PostgreSQL\\16\\bin\\psql.exe",
		"-U", "postgres",
		"-d", "tododb",
		"-f", "C:\\Users\\Lenovo\\GolandProjects\\todo-app\\backend\\db\\migration\\init_schema.sql",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("SQL migration error: %s\n%s", err, output)
	}
	fmt.Println("SQL migration applied successfully")
}

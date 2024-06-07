package setup

import (
	"fmt"
	"log"
	"os"
	"todolist_sprint_asia/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if err := createDatabaseIfNotExists(db, "todolist_sprint_asia"); err != nil {
		panic("failed to create database")
	}

	dsn = fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s sslmode=%s TimeZone=%s",
		host, dbname, user, password, port, sslmode, timezone)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Task{}, &models.SubTask{})

	if err != nil {
		log.Fatalf("failed to migrate tables: %v", err)
	}

	return db
}

func createDatabaseIfNotExists(db *gorm.DB, dbName string) error {
	var count int64
	if err := db.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", dbName).Row().Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		if err := db.Exec("CREATE DATABASE " + dbName); err != nil {
			return nil
		}
	}

	return nil
}

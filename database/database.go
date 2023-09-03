package database

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    models "task-5-pbi-btpns-RoniRagilImanKhoirul/models"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDb() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    fmt.Println("Connection successful!")

    // Perform migrations for User and Photo tables.
    db.AutoMigrate(&models.User{})
    db.AutoMigrate(&models.Photo{})

    DB = db
}

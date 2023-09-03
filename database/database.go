package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	models "task-5-pbi-btpns-RoniRagilImanKhoirul/models"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "rakamin"
)

// ConnectDb menginisialisasi koneksi ke database PostgreSQL dan melakukan migrasi tabel.
func ConnectDb() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi berhasil!")

	// Melakukan migrasi tabel User dan Photo.
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Photo{})

	DB = db
}

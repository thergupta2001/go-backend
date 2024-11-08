package api

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"github.com/thergupta2001/go-backend.git/models"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var JWTSecret string

func SetupDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set in .env")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.AutoMigrate(&models.Doctor{}, &models.Receptionist{}, &models.Patient{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	DB = db
	log.Println("Database connected successfully")
}

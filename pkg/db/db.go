package db

import (

	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Book model that maps to the "books" table
type Book struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string
	Author        string
	PublishedYear int
}

func Connect() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not loaded")
	}

	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")
	DbUsername := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUsername, DbPassword, DbHost, DbPort, DbName)

	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	sqlDB, err := dbConnection.DB()
	if err != nil {
		panic("Failed to get DB handle: " + err.Error())
	}
	if err := sqlDB.Ping(); err != nil {
		panic("Database is unreachable: " + err.Error())
	}

	dbConnection.AutoMigrate(&Book{})
	DB = dbConnection
	return dbConnection
}


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
	// Load .env file (optional) 
	if err := godotenv.Load(); err != nil {
	fmt.Println("Warning: .env file not loaded")
}

	// Read environment variables
	DbHost := os.Getenv("DB_HOST")
	DbName := os.Getenv("DB_NAME")
	DbUsername := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")

	// DSN (Data Source Name) format
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUsername, DbPassword, DbHost, DbName)
		
	fmt.Printf("DB credentials -> host: %s, user: %s, password: %s, dbname: %s\n",
    DbHost, DbUsername, DbPassword, DbName)

	var dbConnection *gorm.DB
	var err error
  
	// Retry logic to wait for database
	for i := 0; i < 10; i++ {
		dbConnection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("DB connected successfully")
			break
		}
	}

	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	dbConnection.AutoMigrate(&Book{})

	DB = dbConnection
	return dbConnection
}

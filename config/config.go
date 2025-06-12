package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MysqlDBName   string
	MysqlUser     string
	MysqlPassword string
	MysqlHost     string
}

func LoadConfig() *Config {
	// Load from .env file (optional in dev)
	if err := godotenv.Load("../../.env"); err != nil {
	fmt.Println("Warning: .env file not loaded")
}


	config := &Config{
		MysqlDBName:   os.Getenv("DB_NAME"),
		MysqlUser:     os.Getenv("DB_USER"),
		MysqlPassword: os.Getenv("DB_PASSWORD"),
		MysqlHost:     os.Getenv("DB_HOST"),
	}

	return config
}

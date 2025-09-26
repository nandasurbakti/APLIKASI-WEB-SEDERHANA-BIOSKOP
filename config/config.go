package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	JWTSecret string
}

func GetDbConfig() DbConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file tidak ditemukan", err)
	}
	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return DbConfig{
		Host:      getEnv("DB_HOST", "default"),
		Port:      port,
		User:      getEnv("DB_USER", "default"),
		Password:  getEnv("DB_PASSWORD", "default"),
		DbName:    getEnv("DB_NAME", "default"),
		JWTSecret: getEnv("JWT_SECRET", "default"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
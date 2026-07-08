package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppEnv  string
	Port    string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBTimeZone string

	JWTSecret        string
	JWTAccessExpire  int
	JWTRefreshExpire int
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName:          getEnv("APP_NAME", "Tetra Server"),
		AppEnv:           getEnv("APP_ENV", "development"),
		Port:             getEnv("PORT", "8080"),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBUser:           getEnv("DB_USER", "postgres"),
		DBPassword:       getEnv("DB_PASSWORD", ""),
		DBName:           getEnv("DB_NAME", "tetra"),
		DBSSLMode:        getEnv("DB_SSLMODE", "disable"),
		DBTimeZone:       getEnv("DB_TIMEZONE", "UTC"),
		JWTSecret:        getEnv("JWT_SECRET", "super-secret-key"),
		JWTAccessExpire:  getEnvInt("JWT_ACCESS_EXPIRE", 15),
		JWTRefreshExpire: getEnvInt("JWT_REFRESH_EXPIRE", 168),
	}

	log.Printf("Config loaded (%s)", cfg.AppEnv)

	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return i
}

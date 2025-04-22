package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
    DB         *gorm.DB
    JWTSecret  string
    AuthExp    int
    RefreshExp int
    AppPort    string
}

func LoadConfig() (*Config, error) {
    godotenv.Load()

    db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return &Config{
        DB:         db,
        JWTSecret:  os.Getenv("JWT_SECRET"),
        AuthExp:    getIntEnv("AUTH_TOKEN_EXP", 3600),
        RefreshExp: getIntEnv("REFRESH_TOKEN_EXP", 604800),
        AppPort:    getEnv("APP_PORT", "3080"),
    }, nil
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}

func getIntEnv(key string, fallback int) int {
    valueStr := getEnv(key, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
        return value
    }
    return fallback
}
package config

import (
	"fmt"
	"os"
)

func DSN() (string, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbDatabase == "" {
		return "", fmt.Errorf("ERROR : required environment variable not found")
	}
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbDatabase,
	) + "parseTime=true&collation=utf8mb4_bin", nil
}

func Port() string {
	return GetEnvOrDefault("PORT", "8080")
}

func GetEnvOrDefault(envPath string, defaultEnv string) string {
	env := os.Getenv(envPath)
	if env == "" {
		return defaultEnv
	}
	return env
}

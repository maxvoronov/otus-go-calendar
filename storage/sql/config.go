package sql

import "os"

// DatabaseConfig struct
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// CreateConfigFromEnvironment Read environment variables into new DatabaseConfig
func CreateConfigFromEnvironment() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
	}
}

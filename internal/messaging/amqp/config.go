package amqp

import "os"

// Config struct
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
}

// CreateConfigFromEnvironment Read environment variables into new Config
func CreateConfigFromEnvironment() *Config {
	return &Config{
		Host:     os.Getenv("MESSAGE_BUS_HOST"),
		Port:     os.Getenv("MESSAGE_BUS_PORT"),
		User:     os.Getenv("MESSAGE_BUS_USER"),
		Password: os.Getenv("MESSAGE_BUS_PASS"),
	}
}

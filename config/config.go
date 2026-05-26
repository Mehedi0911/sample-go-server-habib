package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort    int
	Version     string
	ServiceName string
	JwtSecret   string
}

var config *Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using environment variables", err)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION environment variable not set, using default value '1.0.0'")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HTTP_PORT environment variable not set, using default value '8080'")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVERNAME")
	if serviceName == "" {
		fmt.Println("SERVERNAME environment variable not set, using default value 'SampleServer'")
		os.Exit(1)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("JWT_SECRET environment variable not set, using default value 'your_secret_key_here'")
		os.Exit(1)
	}

	httpPortInt, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Failed to parse HTTP_PORT environment variable")
		os.Exit(1)
	}

	config = &Config{
		Version:     version,
		HTTPPort:    int(httpPortInt), //type casting
		ServiceName: serviceName,
		JwtSecret:   jwtSecret,
	}
}

func GetConfig() *Config {
	if config == nil {
		loadConfig()
	}
	return config
}

package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBUser      string
	DBPassword  string
	DBHost      string
	DBPort      int
	DBName      string
	DB_SSL_Mode string
}

type Config struct {
	HTTPPort    int
	Version     string
	ServiceName string
	JwtSecret   string
	DBConfig    *DBConfig
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

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB_USER environment variable not set")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB_PASSWORD environment variable not set")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB_HOST environment variable not set")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("DB_PORT environment variable not set")
		os.Exit(1)
	}

	dbPortInt, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		fmt.Println("Failed to parse DB_PORT environment variable")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB_NAME environment variable not set")
		os.Exit(1)
	}

	dbSSLMode := os.Getenv("DB_SSL_MODE")
	if dbSSLMode == "" {
		fmt.Println("DB_SSL_MODE environment variable not set")
		os.Exit(1)
	}

	// isSSLMode, err := strconv.ParseBool(dbSSLMode)
	// if err != nil {
	// 	fmt.Println("Failed to parse DB_SSLMODE environment variable")
	// 	os.Exit(1)
	// }

	dbConfig := &DBConfig{
		DBUser:      dbUser,
		DBPassword:  dbPassword,
		DBHost:      dbHost,
		DBPort:      int(dbPortInt),
		DBName:      dbName,
		DB_SSL_Mode: dbSSLMode,
	}

	config = &Config{
		Version:     version,
		HTTPPort:    int(httpPortInt), //type casting
		ServiceName: serviceName,
		JwtSecret:   jwtSecret,
		DBConfig:    dbConfig,
	}
}

func GetConfig() *Config {
	if config == nil {
		loadConfig()
	}
	return config
}

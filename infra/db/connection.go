package db

import (
	"fmt"
	"sample-server/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cfg *config.DBConfig) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%v", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DB_SSL_Mode)
}

func NewConnection(cfg *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cfg)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}

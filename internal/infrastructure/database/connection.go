package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

func LoadDb() *dbConfig {
	return &dbConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		DbName:   os.Getenv("POSTGRES_DB_NAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}
}

func ConnectDb() *gorm.DB {
	cfg := LoadDb()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
			cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port,
		),
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("Erro ao conectar no banco:", err)
		panic(err)
	}

	return db
}

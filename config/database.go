package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CredentialDB struct {
	username string
	password string
	host     string
	port     string
	name     string
}

func InitDB() *gorm.DB {

	var dbUrl = CredentialDB{
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOSTNAME"),
		port:     os.Getenv("DB_PORT"),
		name:     os.Getenv("DB_NAME"),
	}
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUrl.username, dbUrl.password, dbUrl.host, dbUrl.port, dbUrl.name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connect database")
	}

	return db
}

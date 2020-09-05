package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Instance is a exported variable of a db
var Instance *gorm.DB
var env = godotenv.Load()

type connection struct {
	User     string
	Password string
	Host     string
	Port     string
	Schema   string
}

func initDb() *gorm.DB {
	conn := connection{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Schema:   os.Getenv("MYSQL_SCHEMA"),
	}

	credential := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conn.User, conn.Password, conn.Host, conn.Port, conn.Schema)
	db, err := gorm.Open(mysql.Open(credential), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func init() {
	Instance = initDb()
}

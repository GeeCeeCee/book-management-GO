package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func Connect() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOSTNAME")+")/"+os.Getenv("DB_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local"   

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db=d
}

func GetDB ()*gorm.DB {
	return db
}
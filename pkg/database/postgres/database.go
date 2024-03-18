package postgres

import (
	"fmt"
	models2 "github.com/localpurpose/vk-filmoteka/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DBInstance struct {
	DB *gorm.DB
}

var DB DBInstance

func ConnectDB() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err)
	}

	log.Println("Connected to DB")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Runing migrations")
	err = db.AutoMigrate(&models2.Actor{}, &models2.Person{}, &models2.Movie{}, &models2.User{})
	if err != nil {
		log.Println(err)
	}

	DB = DBInstance{
		DB: db,
	}
}

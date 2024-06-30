package database

import (
	"log"
	"os"

	"replica-website-jeketi-backend-go/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB

	env := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AutoMigrate(&model.User{}, &model.Session{}, &model.Product{})
	if err != nil {
		log.Fatal(err)
	}

}

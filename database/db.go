package database

import (
	"fmt"
	"log"
	"mygram/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("PGHOST")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbPort   = os.Getenv("PGPORT")
	dbname   = os.Getenv("PGDATABASE")
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("success connect to database")
	db.Debug().AutoMigrate(models.SocialMedia{}, models.Photo{}, models.Comment{}, models.User{})
}

func GetDB() *gorm.DB {
	return db
}

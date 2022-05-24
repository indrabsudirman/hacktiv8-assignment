package database

import (
	"fmt"
	"hacktiv8-day07/sesi07-gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = "Indra19"
	DB_NAME = "sesi07"
)

func StartDB() *gorm.DB {
	dns := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Default().Println("connected to database successfully")
	err = migration(db)
	if err != nil {
		panic(err)
	}
	return db
}

func migration(db *gorm.DB) error {
	if err := db.AutoMigrate(models.User{}); err != nil {
		return err

	}
	if err := db.AutoMigrate(models.Product{}); err != nil {
		return err

	}
	return nil

}

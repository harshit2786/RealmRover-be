package db

import (
	"fmt"
	"log"
	"realmrovers/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if(err != nil){
		log.Fatalf("Unable to connect to database: %v", err)
	}
	log.Println("Database connected successfully")
	return db
}

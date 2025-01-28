package db

import (
	"log"
	"realmrovers/model"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	
	err := db.AutoMigrate(
		&model.User{},
		&model.Avatar{},
		&model.Realms{},
		&model.Elements{},
		&model.RealmElements{},
	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully")
}

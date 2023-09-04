package helper

import (
	"bwa_startup/campaign"
	"bwa_startup/user"
	"log"

	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	var err error

	err = user.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Error migrating user: %v", err)
	}

	err = campaign.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Error migrating user: %v", err)
	}

	log.Println("Successfully migrated")

}

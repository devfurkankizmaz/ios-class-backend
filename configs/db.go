package configs

import (
	"log"
	"os"

	"github.com/devfurkankizmaz/iosclass-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBConnection() *gorm.DB {
	var err error

	DB, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	query := `
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	`

	if err := DB.Exec(query).Error; err != nil {
		log.Fatal("UUID extension installation failed: ", err.Error())
		os.Exit(1)
	}

	err = DB.AutoMigrate(&models.User{})
	err = DB.AutoMigrate(&models.Gallery{})
	err = DB.AutoMigrate(&models.Visit{})
	err = DB.AutoMigrate(&models.Place{})
	err = DB.AutoMigrate(&models.Address{})

	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Connected to DB")
	return DB
}

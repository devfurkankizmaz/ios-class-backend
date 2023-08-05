package main

import (
	"log"
	"os"

	"github.com/devfurkankizmaz/iosclass-backend/configs"
	"github.com/devfurkankizmaz/iosclass-backend/models"
)

func init() {
	configs.NewDBConnection()
}

func main() {
	err := configs.App().DB.AutoMigrate(&models.User{}, &models.Travel{}, &models.Address{})
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}
}

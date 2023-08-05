package configs

import (
	"gorm.io/gorm"
)

type Application struct {
	DB *gorm.DB
}

func App() Application {
	app := &Application{}
	app.DB = NewDBConnection()
	return *app
}

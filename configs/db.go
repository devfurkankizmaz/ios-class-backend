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

	DB.Commit()

	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	query2 := `
				        ALTER TABLE visits
						ADD CONSTRAINT fk_place
						FOREIGN KEY (place_id)
						REFERENCES places (ID)
						ON DELETE CASCADE;
				    `

	query3 := `ALTER TABLE users
	        ADD CONSTRAINT fk_user_travels FOREIGN KEY (user_id) REFERENCES travels(id) ON DELETE CASCADE;`
	query4 := `CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email ON users(email);`

	DB.Exec(query2)
	DB.Exec(query3)
	DB.Exec(query4)

	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Connected to DB")
	return DB
}

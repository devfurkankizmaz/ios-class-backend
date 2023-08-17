package configs

import (
	"github.com/devfurkankizmaz/iosclass-backend/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBConnection() *gorm.DB {
	var err error

	DB, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Travel{}, &models.Address{})

	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	// Add new fields Latitude and Longitude to Address model
	err = DB.AutoMigrate(&models.Address{
		Latitude:  0.0,
		Longitude: 0.0,
	})

	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	query1 := `ALTER TABLE users
           ADD CONSTRAINT IF NOT EXISTS fk_user_addresses FOREIGN KEY (user_id) REFERENCES addresses(id) ON DELETE CASCADE;`

	query2 := `ALTER TABLE users
           ADD CONSTRAINT IF NOT EXISTS fk_user_travels FOREIGN KEY (user_id) REFERENCES travels(id) ON DELETE CASCADE;`

	DB.Exec(query1)
	DB.Exec(query2)
	query := `CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email ON users(email);`

	DB.Exec(query)
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Connected to DB")
	return DB
}

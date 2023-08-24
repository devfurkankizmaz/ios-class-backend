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

	// Drop and re-create the Address table with new fields

	err = DB.Migrator().DropTable(&models.Gallery{})
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	err = DB.AutoMigrate(&models.Gallery{})

	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	query := `
		        ALTER TABLE galleries
				ADD CONSTRAINT fk_place
				FOREIGN KEY (place_id)
				REFERENCES places (ID)
				ON DELETE CASCADE;
		    `

	//	query2 := `ALTER TABLE users
	//         ADD CONSTRAINT fk_user_travels FOREIGN KEY (user_id) REFERENCES travels(id) ON DELETE CASCADE;`
	//query := `CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email ON users(email);`

	DB.Exec(query)
	//DB.Exec(query2)
	//DB.Exec(query)

	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Connected to DB")
	return DB
}

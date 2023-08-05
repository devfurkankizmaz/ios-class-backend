package configs

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBConnection() *gorm.DB {
	var err error
	dbUrl := os.Getenv("DB_URL")
	println(dbUrl)

	DB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
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

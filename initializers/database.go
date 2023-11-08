package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	// log.Print("Connect TODB INVOKED!")
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// log.Printf("DB: %T", DB)
		log.Fatal("Failed to connect to database")
	} else {
		DB = db
	}

}

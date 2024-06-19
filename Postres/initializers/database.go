package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnecttoDB() {
	var err error
	dsn := "host=localhost user=test password=ha dbname=test port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB ")
	}
	// sqlDB, err := DB.DB()
	// if err != nil {
	// 	log.Fatal("failed to get database instance:", err)
	// }
	// defer sqlDB.Close()

	// // Example query to check the connection
	// err = sqlDB.Ping()
	// if err != nil {
	// 	log.Fatal("failed to ping database:", err)
	// }

	log.Println("Database connection is alive")
}

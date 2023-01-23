package database

import (
	"example/rest-api-fiber/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// connectDb
func ConnectDb() {
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("❌ Failed to connect to database. \n", err)
	}

	log.Println("🔌 Connected!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("🚧 Running Migrations!")
	db.AutoMigrate(&models.Todo{})

	DB = Dbinstance{
		Db: db,
	}
}

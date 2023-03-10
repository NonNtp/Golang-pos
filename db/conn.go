package db

import (
	"go-api/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_DNS")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	Conn = db
}

func Migrate() {
	Conn.AutoMigrate(
		&model.Category{},
		&model.Product{},
		&model.Order{},
		&model.OrderItem{},
	)
}

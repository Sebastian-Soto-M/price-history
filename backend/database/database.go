package database

import (
	"github.com/Sebastian-Soto-M/price-history/exceptions"
	"github.com/Sebastian-Soto-M/price-history/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DBNAME = "price-history.db"

var dbPanic = exceptions.Panic("Database")

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DBNAME), &gorm.Config{})
	dbPanic(err)
	return db
}

func Migrate() {
	db := Connect()
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Store{})
}

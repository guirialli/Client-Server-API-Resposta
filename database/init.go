package database

import (
	"log"

	"github.com/guirialli/dolar-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("contabil.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	log.Println("Sucessfully initialized database!")
}

func Migrate() {
	if DB == nil {
		Init()
	}
	DB.AutoMigrate(&models.Contabil{})
	log.Println("Sucessfully migrated database!")
}

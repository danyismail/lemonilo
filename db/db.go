package db

import (
	"lemonilo/model"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func New() *gorm.DB {
	dbDriver := os.Getenv("DB_DRIVER")
	dbURL := os.Getenv("DB_URL")
	if "" == dbDriver || "" == dbURL {
		log.Fatalln("Database credentials not define.")
	}
	db, err := gorm.Open(dbDriver, dbURL)
	if err != nil {
		log.Fatalln(err)
	}

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)

	var user model.User
	db.AutoMigrate(&user)
	return db
}

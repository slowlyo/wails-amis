package sqlite

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"wails-amis/backend/models"
	"wails-amis/backend/pkg/file"
)

var dbFile = "./data/app.db"

func Connect() *gorm.DB {
	file.InitFile(dbFile)

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	err = db.AutoMigrate(&models.Settings{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return db
}

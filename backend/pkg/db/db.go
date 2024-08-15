package db

import (
	"gorm.io/gorm"
	"wails-amis/backend/pkg/db/sqlite"
)

var Instance *gorm.DB

func Connect() {
	Instance = sqlite.Connect()
}

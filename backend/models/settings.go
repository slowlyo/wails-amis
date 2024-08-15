package models

import "gorm.io/gorm"

type Settings struct {
	gorm.Model
	Key   string `gorm:"unique"`
	Value string
}

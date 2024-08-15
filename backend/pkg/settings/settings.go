package settings

import (
	"gorm.io/gorm/clause"
	"wails-amis/backend/models"
	"wails-amis/backend/pkg/db"
)

func Get(key string) string {
	var settings models.Settings

	if err := db.Instance.Where("key = ?", key).First(&settings).Error; err != nil {
		return ""
	}

	return settings.Value
}

func Set(key, value string) bool {
	settings := &models.Settings{
		Key:   key,
		Value: value,
	}

	result := db.Instance.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"value": value}),
	}).Create(settings)

	return result.Error == nil
}

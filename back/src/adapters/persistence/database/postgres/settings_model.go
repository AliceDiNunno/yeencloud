package postgres

import (
	"gorm.io/gorm"
)

type Settings struct {
	gorm.Model

	Key   string
	Value string
}

func (db *Database) GetSettingsValue(key string) string {
	var settings Settings

	db.engine.Where("key = ?", key).First(&settings)

	return settings.Value
}

func (db *Database) SetSettingsValue(key string, value string) {

}

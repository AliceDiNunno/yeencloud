package postgres

import "gorm.io/gorm"

type CloudObject struct {
	gorm.Model

	ID   string `gorm:"type:uuid;primary_key"`
	Name string
}

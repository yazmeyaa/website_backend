package model

import "gorm.io/gorm"

type StaticFile struct {
	gorm.Model
	FileName string `gorm:"uniqueIndex"`
}

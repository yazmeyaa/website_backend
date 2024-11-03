package model

import "gorm.io/gorm"

type StaticFile struct {
	gorm.Model
	Name     string
	FileName string `gorm:"uniqueIndex"`
}

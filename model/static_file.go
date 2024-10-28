package model

import "gorm.io/gorm"

type StaticFile struct {
	gorm.Model
	Path string `gorm:"uniqueIndex"`
}

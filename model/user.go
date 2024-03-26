package model

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"type: varchar(64);unique"`
	Password string `gorm:"type: varchar(255)"`
}

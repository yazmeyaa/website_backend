package model

type Project struct {
	ID          int     `gorm:"primaryKey"`
	Name        string  `gorm:"type: varchar(255)"`
	Description string  `gorm:"type: varchar(255)"`
	Href        *string `gorm:"type: varchar(255)"`
	Img         string  `gorm:"type: varchar(255)"`
	GithubUrl   *string `gorm:"type: varchar(255)"`
	ImgUrl      *string `gorm:"type: string"`
}

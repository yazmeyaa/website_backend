package model

type Project struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"type: varchar(255)"`
	Description string
	Href        *string
	Img         string
	GithubUrl   *string
}

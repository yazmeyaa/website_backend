package model

import "yazmeyaa_projects/data/response"

type Project struct {
	ID          int     `gorm:"primaryKey"`
	Name        string  `gorm:"type: varchar(255)"`
	Description string  `gorm:"type: varchar(255)"`
	Href        *string `gorm:"type: varchar(255)"`
	Img         string  `gorm:"type: varchar(255)"`
	GithubUrl   *string `gorm:"type: varchar(255)"`
	ImgUrl      *string `gorm:"type: string"`
}

func (p *Project) ToHTTPResponse() response.ProjectsResponse {
	return response.ProjectsResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Href:        p.Href,
		Img:         p.Img,
		GithubUrl:   p.GithubUrl,
		ImgUrl:      p.ImgUrl,
	}
}

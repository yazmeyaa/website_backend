package repository

import "yazmeyaa_projects/model"

type StaticFileRepository interface {
	Create(path string) (*model.StaticFile, error)
	GetById(id int) (*model.StaticFile, error)
	Delete(id int) error
}

package repository

import "yazmeyaa_projects/model"

type StaticFileRepository interface {
	Create(fileName string) (*model.StaticFile, error)
	GetByFileName(fileName string) (*model.StaticFile, error)
	GetById(id int) (*model.StaticFile, error)
	Delete(id int) error
}

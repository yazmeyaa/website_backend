package repository

import (
	"errors"
	"yazmeyaa_projects/helper"
	"yazmeyaa_projects/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{Db: Db}
}

func (repository *UserRepositoryImpl) Create(user model.User) {
	result := repository.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

func (repository *UserRepositoryImpl) FindByUsername(username string) (model.User, error) {
	var user model.User
	result := repository.Db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return user, errors.New("user not found")
	} else {
		return user, nil
	}
}

package repository

import (
	"errors"
	"yazmeyaa_projects/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (repository *UserRepositoryImpl) Create(data CreateUserData) (*model.User, error) {
	var user model.User
	user.Username = data.Username
	user.Password = data.Password

	if err := repository.Db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) FindByUsername(username string) (*model.User, error) {
	var user model.User
	result := repository.Db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

package repository

import "yazmeyaa_projects/model"

type CreateUserData struct {
	Username string
	Password string
}

type UserRepository interface {
	FindByUsername(username string) (user *model.User, err error)
	Create(user CreateUserData) (*model.User, error)
	//Update(user model.User)
	//Delete(user model.User)
}

package repository

import "yazmeyaa_projects/model"

type UserRepository interface {
	FindByUsername(username string) (user model.User, err error)
	Create(user model.User)
	//Update(user model.User)
	//Delete(user model.User)
}

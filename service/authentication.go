package service

import (
	"yazmeyaa_projects/data/request"
	"yazmeyaa_projects/model"
)

type AuthService interface {
	CheckAuth(credentails request.AuthCredentails) (user *model.User, err error)
	Create(credentails request.AuthCredentails) (user *model.User, err error)
}

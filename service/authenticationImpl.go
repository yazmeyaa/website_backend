package service

import (
	"errors"
	"fmt"
	"yazmeyaa_projects/data/request"
	"yazmeyaa_projects/helper"
	"yazmeyaa_projects/model"
	"yazmeyaa_projects/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewAuthService(repository repository.UserRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: repository,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) GetUser(username string) (model.User, error) {
	user, err := service.UserRepository.FindByUsername(username)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (service *AuthServiceImpl) CheckAuth(credentails request.AuthCredentails) (user model.User, err error) {
	validationError := service.Validate.Struct(credentails)
	if validationError != nil {
		return user, errors.New("not valid json")
	}
	result, err := service.GetUser(credentails.Username)
	fmt.Println(string(result.Password))

	if err != nil {
		return result, err
	}

	compareError := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(credentails.Password))
	if compareError != nil {
		return result, errors.New("wrong password")
	}

	return result, nil
}

func (service *AuthServiceImpl) Create(credentails request.AuthCredentails) (model.User, error) {
	validationError := service.Validate.Struct(credentails)
	if validationError != nil {
		return model.User{}, validationError
	}

	_, existUserError := service.GetUser(credentails.Username)
	if existUserError == nil {
		return model.User{}, errors.New("username is already taken")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(credentails.Password), 5)
	helper.ErrorPanic(err)

	user := model.User{
		Password: string(hash),
		Username: credentails.Username,
	}

	service.UserRepository.Create(user)
	return user, nil
}

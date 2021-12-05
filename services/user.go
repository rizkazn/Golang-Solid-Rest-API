package services

import (
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/interfaces"
	"github.com/rizkazn/gorestfull/models"
)

type user_service struct {
	rp interfaces.UserRepository
}

func UserService(rps interfaces.UserRepository) *user_service {
	return &user_service{rps}
}

func (us *user_service) GetAll() (*helpers.Respond, error) {
	data, err := us.rp.FindAll()

	if err != nil {
		return nil, err

	}

	respon := helpers.Respond{
		Status:  200,
		IsError: false,
		Data:    data,
	}

	return &respon, nil
}

func (us *user_service) Add(user *models.User) (*helpers.Respond, error) {
	err := us.rp.Save(user)

	if err != nil {
		return nil, err

	}

	respon := helpers.Respond{
		Status:  200,
		IsError: false,
		Data:    "The user has been successfully saved",
	}

	return &respon, nil
}

func (us *user_service) GetPassword(username string) (string, error) {
	data, err := us.rp.GetPassword(username)

	if err != nil {
		return "", err
	}
	return data, nil
}

func (us *user_service) GetUsername(username string) (string, error) {
	data, err := us.rp.GetUsername(username)

	if err != nil {
		return "", err
	}
	return data, nil
}

func (us *user_service) GetEmail(username string) (string, error) {
	data, err := us.rp.GetEmail(username)

	if err != nil {
		return "", err
	}
	return data, nil
}

func (us *user_service) GetRole(username string) (string, error) {
	data, err := us.rp.GetRole(username)

	if err != nil {
		return "", err
	}
	return data, nil
}

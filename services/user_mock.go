package services

import (
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/models"
	"github.com/stretchr/testify/mock"
)

type UserMockServ struct {
	Mock mock.Mock
}

var dataUser = helpers.Respond{
	Status: 200,
}

func (mock *UserMockServ) GetAll() (*helpers.Respond, error) {
	args := mock.Mock.Called()

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *UserMockServ) Add(user *models.User) (*helpers.Respond, error) {
	args := mock.Mock.Called(user)

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *UserMockServ) Update(user *models.User, id string) (*helpers.Respond, error) {
	args := mock.Mock.Called(user, id)

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *UserMockServ) Delete(id string) (*helpers.Respond, error) {
	args := mock.Mock.Called(id)

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *UserMockServ) GetPassword(username string) (string, error) {
	args := mock.Mock.Called(username)

	if args.Get(0) == nil {
		data.IsError = true
		return "", nil
	}

	data.IsError = false
	return "", nil
}
func (mock *UserMockServ) GetUsername(username string) (string, error) {
	args := mock.Mock.Called(username)

	if args.Get(0) == nil {
		data.IsError = true
		return "", nil
	}

	data.IsError = false
	return "", nil
}

func (mock *UserMockServ) GetEmail(username string) (string, error) {
	args := mock.Mock.Called(username)

	if args.Get(0) == nil {
		data.IsError = true
		return "", nil
	}

	data.IsError = false
	return "", nil
}

func (mock *UserMockServ) GetRole(username string) (string, error) {
	args := mock.Mock.Called(username)

	if args.Get(0) == nil {
		data.IsError = true
		return "", nil
	}

	data.IsError = false
	return "", nil
}

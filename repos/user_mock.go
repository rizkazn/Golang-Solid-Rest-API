package repos

import (
	"time"

	"github.com/rizkazn/gorestfull/models"
	"github.com/stretchr/testify/mock"
)

type UserMockRepos struct {
	Mock mock.Mock
}

var dataUser = models.User{
	Id:        1,
	Name:      "Name for Testing",
	Username:  "Username for Testing",
	Email:     "Email for Testing",
	Password:  "Password for Testing",
	Role:      "Role for Testing",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var datumUser = models.Users{
	dataUser,
}

func (mock *UserMockRepos) FindAll() (*models.Users, error) {
	// args := mock.Mock.Called()
	// result := args.Get(0)
	// return result.(*models.Users), args.Error(1)
	args := mock.Mock.Called()

	if args == nil {
		return nil, nil
	}

	return &datumUser, nil
}

func (mock *UserMockRepos) Save(user *models.User) error {
	args := mock.Mock.Called(user)

	if args == nil {
		return nil
	}

	return nil
}

func (mock *UserMockRepos) GetPassword(username string) (string, error) {
	args := mock.Mock.Called(username)

	if args == nil {
		return "", nil
	}

	return "", nil
}

func (mock *UserMockRepos) GetUsername(username string) (string, error) {
	args := mock.Mock.Called(username)

	if args == nil {
		return "", nil
	}

	return "", nil
}

func (mock *UserMockRepos) GetEmail(username string) (string, error) {
	args := mock.Mock.Called(username)

	if args == nil {
		return "", nil
	}

	return "", nil
}

func (mock *UserMockRepos) GetRole(username string) (string, error) {
	args := mock.Mock.Called(username)

	if args == nil {
		return "", nil
	}

	return "", nil
}

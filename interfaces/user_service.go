package interfaces

import (
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/models"
)

type UserServices interface {
	GetAll() (*helpers.Respond, error)
	Add(user *models.User) (*helpers.Respond, error)
	GetPassword(username string) (string, error)
	GetUsername(username string) (string, error)
	GetEmail(username string) (string, error)
	GetRole(username string) (string, error)
}

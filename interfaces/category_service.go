package interfaces

import (
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/models"
)

type CategoryServices interface {
	GetAll() (*helpers.Respond, error)
	Add(category *models.Category) (*helpers.Respond, error)
	Update(category *models.Category, id string) (*helpers.Respond, error)
	Delete(id string) (*helpers.Respond, error)
}

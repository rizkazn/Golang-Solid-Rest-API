package interfaces

import (
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/models"
)

type ProductServices interface {
	GetAll() (*helpers.Respond, error)
	Add(product *models.Product) (*helpers.Respond, error)
	Update(product *models.Product, id string) (*helpers.Respond, error)
	Delete(id string) (*helpers.Respond, error)
	SearchByName(name string) (*helpers.Respond, error)
	SortByName() (*helpers.Respond, error)
	SortByCat() (*helpers.Respond, error)
	SortByNewest() (*helpers.Respond, error)
	SortByPrice() (*helpers.Respond, error)
}

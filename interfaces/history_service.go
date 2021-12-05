package interfaces

import (
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/models"
)

type HistoryServices interface {
	GetAll() (*helpers.Respond, error)
	Add(history *models.History) (*helpers.Respond, error)
	Update(history *models.History, id string) (*helpers.Respond, error)
	Delete(id string) (*helpers.Respond, error)
}

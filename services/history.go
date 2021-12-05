package services

import (
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/interfaces"
	"github.com/rizkazn/gorestfull/models"
)

type history_service struct {
	rp interfaces.HistRepository
}

func HistoryService(rps interfaces.HistRepository) *history_service {
	return &history_service{rps}
}

func (hist *history_service) GetAll() (*helpers.Respond, error) {
	data, err := hist.rp.FindAll()

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

func (hist *history_service) Add(history *models.History) (*helpers.Respond, error) {
	err := hist.rp.Save(history)

	if err != nil {
		return nil, err

	}

	respon := helpers.Respond{
		Status:  200,
		IsError: false,
		Data:    "The history has been successfully saved",
	}

	return &respon, nil
}

func (hist *history_service) Update(history *models.History, id string) (*helpers.Respond, error) {
	err := hist.rp.Update(history, id)

	if err != nil {
		return nil, err

	}

	respon := helpers.Respond{
		Status:  200,
		IsError: false,
		Data:    "The history has been successfully updated",
	}

	return &respon, nil
}

func (hist *history_service) Delete(id string) (*helpers.Respond, error) {
	err := hist.rp.Remove(id)

	if err != nil {
		return nil, err

	}

	respon := helpers.Respond{
		Status:  200,
		IsError: false,
		Data:    "The history has been successfully deleted",
	}

	return &respon, nil
}

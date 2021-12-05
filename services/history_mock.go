package services

import (
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/models"
	"github.com/stretchr/testify/mock"
)

type HistoryMockServ struct {
	Mock mock.Mock
}

var dataHist = helpers.Respond{
	Status: 200,
}

func (mock *HistoryMockServ) GetAll() (*helpers.Respond, error) {
	args := mock.Mock.Called()

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *HistoryMockServ) Add(history *models.History) (*helpers.Respond, error) {
	args := mock.Mock.Called(history)

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *HistoryMockServ) Update(history *models.History, id string) (*helpers.Respond, error) {
	args := mock.Mock.Called(history, id)

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *HistoryMockServ) Delete(id string) (*helpers.Respond, error) {
	args := mock.Mock.Called(id)

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

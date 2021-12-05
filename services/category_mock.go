package services

import (
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/models"
	"github.com/stretchr/testify/mock"
)

type CategoryMockServ struct {
	Mock mock.Mock
}

var dataCat = helpers.Respond{
	Status: 200,
}

func (mock *CategoryMockServ) GetAll() (*helpers.Respond, error) {
	args := mock.Mock.Called()

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *CategoryMockServ) Add(category *models.Category) (*helpers.Respond, error) {
	args := mock.Mock.Called(category)

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *CategoryMockServ) Update(category *models.Category, id string) (*helpers.Respond, error) {
	args := mock.Mock.Called(category, id)

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *CategoryMockServ) Delete(id string) (*helpers.Respond, error) {
	args := mock.Mock.Called(id)

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

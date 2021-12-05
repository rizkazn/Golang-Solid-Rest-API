package services

import (
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/models"
	"github.com/stretchr/testify/mock"
)

type ProductMockServ struct {
	Mock mock.Mock
}

var data = helpers.Respond{
	Status: 200,
}

func (mock *ProductMockServ) GetAll() (*helpers.Respond, error) {
	args := mock.Mock.Called()

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *ProductMockServ) Add(product *models.Product) (*helpers.Respond, error) {
	args := mock.Mock.Called(product)

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *ProductMockServ) Update(product *models.Product, id string) (*helpers.Respond, error) {
	args := mock.Mock.Called(product, id)

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *ProductMockServ) Delete(id string) (*helpers.Respond, error) {
	args := mock.Mock.Called(id)

	if args == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *ProductMockServ) SearchByName(name string) (*helpers.Respond, error) {
	args := mock.Mock.Called(name)

	if args.Get(0) == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *ProductMockServ) SortByName() (*helpers.Respond, error) {
	args := mock.Mock.Called()

	if args.Get(0) == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *ProductMockServ) SortByCat() (*helpers.Respond, error) {
	args := mock.Mock.Called()

	if args.Get(0) == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *ProductMockServ) SortByNewest() (*helpers.Respond, error) {
	args := mock.Mock.Called()

	if args.Get(0) == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

func (mock *ProductMockServ) SortByPrice() (*helpers.Respond, error) {
	args := mock.Mock.Called()

	if args.Get(0) == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}

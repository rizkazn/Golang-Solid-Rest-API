package services

import (
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/interfaces"
	"github.com/rizkazn/gorestfull/models"
)

type product_service struct {
	rp interfaces.ProdRepository
}

func ProductService(rps interfaces.ProdRepository) *product_service {
	return &product_service{rps}
}

func (prod *product_service) GetAll() (*helpers.Respond, error) {
	data, err := prod.rp.FindAll()

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

func (prod *product_service) Add(product *models.Product) (*helpers.Respond, error) {
	err := prod.rp.Save(product)

	if err != nil {
		return nil, err

	}

	respon := helpers.Respond{
		Status:  200,
		IsError: false,
		Data:    "The product has been successfully saved",
	}

	return &respon, nil
}

func (prod *product_service) Update(product *models.Product, id string) (*helpers.Respond, error) {
	err := prod.rp.Update(product, id)

	if err != nil {
		return nil, err

	}

	respon := helpers.Respond{
		Status:  200,
		IsError: false,
		Data:    "The product has been successfully updated",
	}

	return &respon, nil
}

func (prod *product_service) Delete(id string) (*helpers.Respond, error) {
	err := prod.rp.Remove(id)

	if err != nil {
		return nil, err

	}

	respon := helpers.Respond{
		Status:  200,
		IsError: false,
		Data:    "The product has been successfully deleted",
	}

	return &respon, nil
}

func (prod *product_service) SearchByName(name string) (*helpers.Respond, error) {
	data, err := prod.rp.SearchProductBytName(name)

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

func (prod *product_service) SortByName() (*helpers.Respond, error) {
	data, err := prod.rp.OrderProdByName()

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

func (prod *product_service) SortByCat() (*helpers.Respond, error) {
	data, err := prod.rp.OrderProdByCategory()

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

func (prod *product_service) SortByNewest() (*helpers.Respond, error) {
	data, err := prod.rp.OrderProdByNewest()

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

func (prod *product_service) SortByPrice() (*helpers.Respond, error) {
	data, err := prod.rp.OrderProdByPrice()

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

package services

import (
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/interfaces"
	"github.com/rizkazn/gorestfull/models"
)

type category_service struct {
	rp interfaces.CatRepository
}

func CategoryService(rps interfaces.CatRepository) *category_service {
	return &category_service{rps}
}

func (cat *category_service) GetAll() (*helpers.Respond, error) {
	data, err := cat.rp.FindAll()

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

func (cat *category_service) Add(category *models.Category) (*helpers.Respond, error) {
	err := cat.rp.Save(category)

	if err != nil {
		return nil, err

	}

	respon := helpers.Respond{
		Status:  200,
		IsError: false,
		Data:    "The category has been successfully saved",
	}

	return &respon, nil
}

func (cat *category_service) Update(category *models.Category, id string) (*helpers.Respond, error) {
	err := cat.rp.Update(category, id)

	if err != nil {
		return nil, err

	}

	respon := helpers.Respond{
		Status:  200,
		IsError: false,
		Data:    "The category has been successfully updated",
	}

	return &respon, nil
}

func (cat *category_service) Delete(id string) (*helpers.Respond, error) {
	err := cat.rp.Remove(id)

	if err != nil {
		return nil, err

	}

	respon := helpers.Respond{
		Status:  200,
		IsError: false,
		Data:    "The category has been successfully deleted",
	}

	return &respon, nil
}

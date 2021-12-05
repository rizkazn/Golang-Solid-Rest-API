package repos

import (
	"time"

	"github.com/rizkazn/gorestfull/models"
	"github.com/stretchr/testify/mock"
)

type CategoryMockRepos struct {
	Mock mock.Mock
}

var dataCat = models.Category{
	Id:             1,
	Category_name:  "Category Name for testing",
	Category_image: "Category Image for testing",
	CreatedAt:      time.Now(),
	UpdatedAt:      time.Now(),
}

var datumCat = models.Categories{
	dataCat,
}

func (mock *CategoryMockRepos) FindAll() (*models.Categories, error) {
	// args := mock.Mock.Called()
	// result := args.Get(0)
	// return result.(*models.Categories), args.Error(1)
	args := mock.Mock.Called()

	if args == nil {
		return nil, nil
	}

	return &datumCat, nil
}

func (mock *CategoryMockRepos) Save(category *models.Category) error {
	args := mock.Mock.Called(category)

	if args == nil {
		return nil
	}

	return nil
}

func (mock *CategoryMockRepos) Update(category *models.Category, id string) error {
	args := mock.Mock.Called(category, id)

	if args == nil {
		return nil
	}

	return nil
}

func (mock *CategoryMockRepos) Remove(id string) error {
	args := mock.Mock.Called(id)

	if args == nil {
		return nil
	}

	return nil
}

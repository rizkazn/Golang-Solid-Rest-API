package repos

import (
	"time"

	"github.com/rizkazn/gorestfull/models"
	"github.com/stretchr/testify/mock"
)

type ProductMockRepos struct {
	Mock mock.Mock
}

var dataProd = models.Product{
	Id:            1,
	Product_name:  "Product Name for Testing",
	Price:         120000,
	Product_image: "Product Image for testing",
	Category_id:   1,
	CreatedAt:     time.Now(),
	UpdatedAt:     time.Now(),
}

var datumProd = models.Products{
	dataProd,
}

func (mock *ProductMockRepos) FindAll() (*models.Products, error) {
	// args := mock.Mock.Called()
	// result := args.Get(0)
	// return result.(*models.Products), args.Error(1)
	args := mock.Mock.Called()

	if args == nil {
		return nil, nil
	}

	return &datumProd, nil
}

func (mock *ProductMockRepos) Save(product *models.Product) error {
	args := mock.Mock.Called(product)

	if args == nil {
		return nil
	}

	return nil
}

func (mock *ProductMockRepos) Update(product *models.Product, id string) error {
	args := mock.Mock.Called(product, id)

	if args == nil {
		return nil
	}

	return nil
}

func (mock *ProductMockRepos) Remove(id string) error {
	args := mock.Mock.Called(id)

	if args == nil {
		return nil
	}

	return nil
}

func (mock *ProductMockRepos) SearchProductBytName(name string) (*models.Products, error) {
	args := mock.Mock.Called()

	if args.Get(0) == nil {
		return nil, nil
	}

	return nil, nil
}

func (mock *ProductMockRepos) OrderProdByName() (*models.Products, error) {
	args := mock.Mock.Called()

	if args.Get(0) == nil {
		return nil, nil
	}

	return nil, nil
}

func (mock *ProductMockRepos) OrderProdByCategory() (*models.Products, error) {
	args := mock.Mock.Called()

	if args.Get(0) == nil {
		return nil, nil
	}

	return nil, nil
}
func (mock *ProductMockRepos) OrderProdByNewest() (*models.Products, error) {
	args := mock.Mock.Called()

	if args.Get(0) == nil {
		return nil, nil
	}

	return nil, nil
}
func (mock *ProductMockRepos) OrderProdByPrice() (*models.Products, error) {
	args := mock.Mock.Called()

	if args.Get(0) == nil {
		return nil, nil
	}

	return nil, nil
}

// func (mock *ProductMockRepos) Save(product *models.Product) error {
// 	args := mock.Mock.Called(product)
// 	args.Get(0)
// 	return args.Error(1)
// }

// func (mock *ProductMockRepos) Update(product *models.Product, id string) error {
// 	args := mock.Mock.Called(product)
// 	args.Get(0)
// 	return args.Error(1)
// }

// func (mock *ProductMockRepos) Remove(id string) error {
// 	args := mock.Mock.Called(id)
// 	args.Get(0)
// 	return args.Error(1)
// }

// func (mock *ProductMockRepos) SearchProductBytName(name string) (*models.Products, error) {
// 	args := mock.Mock.Called(name)
// 	result := args.Get(0)
// 	return result.(*models.Products), args.Error(1)
// }

// func (mock *ProductMockRepos) OrderProdByName() (*models.Products, error) {
// 	args := mock.Mock.Called()
// 	result := args.Get(0)
// 	return result.(*models.Products), args.Error(1)
// }

// func (mock *ProductMockRepos) OrderProdByCategory() (*models.Products, error) {
// 	args := mock.Mock.Called()
// 	result := args.Get(0)
// 	return result.(*models.Products), args.Error(1)
// }

// func (mock *ProductMockRepos) OrderProdByNewest() (*models.Products, error) {
// 	args := mock.Mock.Called()
// 	result := args.Get(0)
// 	return result.(*models.Products), args.Error(1)
// }

// func (mock *ProductMockRepos) OrderProdByPrice() (*models.Products, error) {
// 	args := mock.Mock.Called()
// 	result := args.Get(0)
// 	return result.(*models.Products), args.Error(1)
// }

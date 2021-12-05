package services

import (
	"testing"
	"time"

	"github.com/rizkazn/gorestfull/models"
	"github.com/rizkazn/gorestfull/repos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepos = repos.ProductMockRepos{Mock: mock.Mock{}}
var productService = product_service{rp: &productRepos}

func TestGetAllData(t *testing.T) {
	t.Run("Get All Product Data", func(t *testing.T) {
		productRepos.Mock.On("FindAll")

		data, err := productService.GetAll()
		result := data.Status

		assert.Equal(t, 200, result, "Status code must be 200")
		assert.Nil(t, err)
	})
}

func TestAddData(t *testing.T) {
	t.Run("Add Product Data", func(t *testing.T) {
		product := models.Product{
			Id:            1,
			Product_name:  "test",
			Price:         120000,
			Product_image: "test",
			Category_id:   1,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		productRepos.Mock.On("Save", &product)

		data, err := productService.Add(&product)
		result := data.Data

		assert.Equal(t, 200, data.Status, "Status code must be 200")
		assert.Equal(t, "The product has been successfully saved", result, "Data must be 'The product has been successfully saved'")
		assert.Nil(t, err)
	})
}

func TestUpdateData(t *testing.T) {
	t.Run("Update Product Data", func(t *testing.T) {
		product := models.Product{
			Id:            1,
			Product_name:  "test",
			Price:         120000,
			Product_image: "test",
			Category_id:   1,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		productRepos.Mock.On("Update", &product, "1")

		data, err := productService.Update(&product, "1")
		result := data.Data

		assert.Equal(t, 200, data.Status, "Status code must be 200")
		assert.Equal(t, "The product has been successfully updated", result, "Data must be 'The product has been successfully updated'")
		assert.Nil(t, err)
	})
}

func TestRemoveData(t *testing.T) {
	t.Run("Remove Product Data", func(t *testing.T) {
		productRepos.Mock.On("Remove", "1")

		data, err := productService.Delete("1")
		result := data.Data

		assert.Equal(t, 200, data.Status, "Status code must be 200")
		assert.Equal(t, "The product has been successfully deleted", result, "Data must be 'The product has been successfully deleted'")
		assert.Nil(t, err)
	})
}

// func TestSearchData(t *testing.T) {
// 	t.Run("Search Product Data by Name", func(t *testing.T) {
// 		productRepos.Mock.On("SearchProductBytName", "test")

// 		data, err := productService.SearchByName("test")
// 		result := data.Data

// 		assert.Equal(t, "test", result, "Data must be 'test'")
// 		assert.Nil(t, err)
// 	})
// }

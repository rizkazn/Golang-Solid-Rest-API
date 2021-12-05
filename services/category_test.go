package services

import (
	"testing"
	"time"

	"github.com/rizkazn/gorestfull/models"
	"github.com/rizkazn/gorestfull/repos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepos = repos.CategoryMockRepos{Mock: mock.Mock{}}
var categoryService = category_service{rp: &categoryRepos}

func TestGetAllDataCat(t *testing.T) {
	t.Run("Get All Category Data", func(t *testing.T) {
		categoryRepos.Mock.On("FindAll")

		data, err := categoryService.GetAll()
		result := data.Status

		assert.Equal(t, 200, result, "Status code must be 200")
		assert.Nil(t, err)
	})
}

func TestAddDataCat(t *testing.T) {
	t.Run("Add Category Data", func(t *testing.T) {
		category := models.Category{
			Id:             1,
			Category_name:  "test",
			Category_image: "test",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		categoryRepos.Mock.On("Save", &category)

		data, err := categoryService.Add(&category)
		result := data.Data

		assert.Equal(t, "The category has been successfully saved", result, "Data must be 'The category has been successfully saved'")
		assert.Equal(t, 200, data.Status, "Status code must be 200")
		assert.Nil(t, err)
	})
}

func TestUpdateDataCat(t *testing.T) {
	t.Run("Update Category Data", func(t *testing.T) {
		category := models.Category{
			Id:             1,
			Category_name:  "test",
			Category_image: "test",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		categoryRepos.Mock.On("Update", &category, "1")

		data, err := categoryService.Update(&category, "1")
		result := data.Data

		assert.Equal(t, 200, data.Status, "Status code must be 200")
		assert.Equal(t, "The category has been successfully updated", result, "Data must be 'The category has been successfully updated'")
		assert.Nil(t, err)
	})
}

func TestRemoveDataCat(t *testing.T) {
	t.Run("Remove Category Data", func(t *testing.T) {
		categoryRepos.Mock.On("Remove", "1")

		data, err := categoryService.Delete("1")
		result := data.Data

		assert.Equal(t, 200, data.Status, "Status code must be 200")
		assert.Equal(t, "The category has been successfully deleted", result, "Data must be 'The category has been successfully deleted'")
		assert.Nil(t, err)
	})
}

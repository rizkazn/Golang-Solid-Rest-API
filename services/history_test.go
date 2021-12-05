package services

import (
	"testing"

	"github.com/rizkazn/gorestfull/models"
	"github.com/rizkazn/gorestfull/repos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var historyRepos = repos.HistoryMockRepos{Mock: mock.Mock{}}
var historyService = history_service{rp: &historyRepos}

func TestGetAllDataHist(t *testing.T) {
	t.Run("Get All History Data", func(t *testing.T) {
		historyRepos.Mock.On("FindAll")

		data, err := historyService.GetAll()
		result := data.Status

		assert.Equal(t, 200, result, "Status code must be 200")
		assert.Nil(t, err)
	})
}

func TestAddDataHist(t *testing.T) {
	t.Run("Add History Data", func(t *testing.T) {
		history := models.History{
			Id:         1,
			No_invoice: "test",
			Cashier:    "test",
			Date:       "test",
			Orders:     1,
			Amount:     1,
		}

		historyRepos.Mock.On("Save", &history)

		data, err := historyService.Add(&history)
		result := data.Data

		assert.Equal(t, "The history has been successfully saved", result, "Data must be 'The history has been successfully saved'")
		assert.Equal(t, 200, data.Status, "Status code must be 200")
		assert.Nil(t, err)
	})
}

func TestUpdateDataHist(t *testing.T) {
	t.Run("Update History Data", func(t *testing.T) {
		history := models.History{
			Id:         1,
			No_invoice: "test",
			Cashier:    "test",
			Date:       "test",
			Orders:     1,
			Amount:     1,
		}

		historyRepos.Mock.On("Update", &history, "1")

		data, err := historyService.Update(&history, "1")
		result := data.Data

		assert.Equal(t, 200, data.Status, "Status code must be 200")
		assert.Equal(t, "The history has been successfully updated", result, "Data must be 'The history has been successfully updated'")
		assert.Nil(t, err)
	})
}

func TestRemoveDataHist(t *testing.T) {
	t.Run("Remove History Data", func(t *testing.T) {
		historyRepos.Mock.On("Remove", "1")

		data, err := historyService.Delete("1")
		result := data.Data

		assert.Equal(t, 200, data.Status, "Status code must be 200")
		assert.Equal(t, "The history has been successfully deleted", result, "Data must be 'The history has been successfully deleted'")
		assert.Nil(t, err)
	})
}

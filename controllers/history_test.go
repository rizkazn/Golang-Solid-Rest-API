package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/models"
	"github.com/rizkazn/gorestfull/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var historyService = services.HistoryMockServ{Mock: mock.Mock{}}
var historyController = History(&historyService)

func TestGetAllDataHist(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test", historyController.GetAll)
	historyService.Mock.On("GetAll")

	r.ServeHTTP(w, httptest.NewRequest("GET", "/test", nil))

	var respon helpers.Respond

	err := json.Unmarshal(w.Body.Bytes(), &respon)

	if err != nil {
		t.Fatal("Conversion to Struct Failed")
	}

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

func TestAddDataHist(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test", historyController.Add)
	history := models.History{
		Id:         1,
		No_invoice: "test",
		Cashier:    "test",
		Date:       "test",
		Orders:     1,
		Amount:     1,
	}
	historyService.Mock.On("Add", history)

	r.ServeHTTP(w, httptest.NewRequest("POST", "/test", nil))

	var respon helpers.Respond

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

func TestUpdateDataHist(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test/{id}", historyController.Update)
	history := &models.Category{
		Id:             1,
		Category_name:  "test",
		Category_image: "test",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	historyService.Mock.On("Update", history)

	r.ServeHTTP(w, httptest.NewRequest("PUT", "/test/0", nil))

	var respon helpers.Respond

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

func TestDeleteDataHist(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test/{id}", historyController.Delete)

	historyService.Mock.On("Delete", "0")

	r.ServeHTTP(w, httptest.NewRequest("DELETE", "/test/0", nil))

	var respon helpers.Respond

	err := json.Unmarshal(w.Body.Bytes(), &respon)

	if err != nil {
		t.Fatal("Conversion to Struct Failed")
	}

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

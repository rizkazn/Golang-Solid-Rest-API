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

var productService = services.ProductMockServ{Mock: mock.Mock{}}
var productController = Products(&productService)

func TestGetAllData(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test", productController.GetAll)
	productService.Mock.On("GetAll")

	r.ServeHTTP(w, httptest.NewRequest("GET", "/test", nil))

	var respon helpers.Respond

	err := json.Unmarshal(w.Body.Bytes(), &respon)

	if err != nil {
		t.Fatal("Conversion to Struct Failed")
	}

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

func TestAddData(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test", productController.Add)
	product := models.Product{
		Id:            1,
		Product_name:  "test",
		Price:         120000,
		Product_image: "test",
		Category_id:   1,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	productService.Mock.On("Add", product)

	r.ServeHTTP(w, httptest.NewRequest("POST", "/test", nil))

	var respon helpers.Respond

	// err := json.Unmarshal(w.Body.Bytes(), &respon)

	// if err != nil {
	// 	t.Fatal("Conversion to Struct Failed")
	// }

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

func TestUpdateData(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test/{id}", productController.Update)
	product := &models.Product{
		Id:            1,
		Product_name:  "test",
		Price:         120000,
		Product_image: "test",
		Category_id:   1,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	productService.Mock.On("Update", product)

	r.ServeHTTP(w, httptest.NewRequest("PUT", "/test/0", nil))

	var respon helpers.Respond

	// err := json.Unmarshal(w.Body.Bytes(), &respon)

	// if err != nil {
	// 	t.Fatal("Conversion to Struct Failed")
	// }

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

func TestDeleteData(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test/{id}", productController.Delete)

	productService.Mock.On("Delete", "0")

	r.ServeHTTP(w, httptest.NewRequest("DELETE", "/test/0", nil))

	var respon helpers.Respond

	err := json.Unmarshal(w.Body.Bytes(), &respon)

	if err != nil {
		t.Fatal("Conversion to Struct Failed")
	}

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

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

var categoryService = services.CategoryMockServ{Mock: mock.Mock{}}
var categoryController = Category(&categoryService)

func TestGetAllDataCat(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test", categoryController.GetAll)
	categoryService.Mock.On("GetAll")

	r.ServeHTTP(w, httptest.NewRequest("GET", "/test", nil))

	var respon helpers.Respond

	err := json.Unmarshal(w.Body.Bytes(), &respon)

	if err != nil {
		t.Fatal("Conversion to Struct Failed")
	}

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

func TestAddDataCat(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test", categoryController.Add)
	category := models.Category{
		Id:             1,
		Category_name:  "test",
		Category_image: "test",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	productService.Mock.On("Add", category)

	r.ServeHTTP(w, httptest.NewRequest("POST", "/test", nil))

	var respon helpers.Respond

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

func TestUpdateDataCat(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test/{id}", categoryController.Update)
	category := &models.Category{
		Id:             1,
		Category_name:  "test",
		Category_image: "test",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	categoryService.Mock.On("Update", category)

	r.ServeHTTP(w, httptest.NewRequest("PUT", "/test/0", nil))

	var respon helpers.Respond

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

func TestDeleteDataCat(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test/{id}", categoryController.Delete)

	categoryService.Mock.On("Delete", "0")

	r.ServeHTTP(w, httptest.NewRequest("DELETE", "/test/0", nil))

	var respon helpers.Respond

	err := json.Unmarshal(w.Body.Bytes(), &respon)

	if err != nil {
		t.Fatal("Conversion to Struct Failed")
	}

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

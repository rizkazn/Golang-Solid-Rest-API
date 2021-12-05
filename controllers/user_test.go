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

var userService = services.UserMockServ{Mock: mock.Mock{}}
var userController = User(&userService)

func TestGetAllDataUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test", userController.GetAll)
	userService.Mock.On("GetAll")

	r.ServeHTTP(w, httptest.NewRequest("GET", "/test", nil))

	var respon helpers.Respond

	err := json.Unmarshal(w.Body.Bytes(), &respon)

	if err != nil {
		t.Fatal("Conversion to Struct Failed")
	}

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

func TestAddDataUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test", userController.Add)
	user := models.User{
		Id:        1,
		Name:      "test",
		Username:  "test",
		Email:     "test",
		Password:  "test",
		Role:      "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	userService.Mock.On("Add", user)

	r.ServeHTTP(w, httptest.NewRequest("POST", "/test", nil))

	var respon helpers.Respond

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respon.IsError, "IsError must be false")
}

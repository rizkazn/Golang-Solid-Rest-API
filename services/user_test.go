package services

import (
	"testing"
	"time"

	"github.com/rizkazn/gorestfull/models"
	"github.com/rizkazn/gorestfull/repos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepos = repos.UserMockRepos{Mock: mock.Mock{}}
var userService = user_service{rp: &userRepos}

func TestGetAllDataUser(t *testing.T) {
	t.Run("Get All User Data", func(t *testing.T) {
		userRepos.Mock.On("FindAll")

		data, err := userService.GetAll()
		result := data.Status

		assert.Equal(t, 200, result, "Status code must be 200")
		assert.Nil(t, err)
	})
}

func TestAddDataUser(t *testing.T) {
	t.Run("Add User Data", func(t *testing.T) {
		user := models.User{
			Id:        1,
			Name:      "test",
			Username:  "test",
			Email:     "test@gmail.com",
			Password:  "test1234",
			Role:      "test",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		userRepos.Mock.On("Save", &user)

		data, err := userService.Add(&user)
		result := data.Data

		assert.Equal(t, "The user has been successfully saved", result, "Data must be 'The user has been successfully saved'")
		assert.Equal(t, 200, data.Status, "Status code must be 200")
		assert.Nil(t, err)
	})
}

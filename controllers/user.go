package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rizkazn/gorestfull/helpers"
	"github.com/rizkazn/gorestfull/interfaces"
	"github.com/rizkazn/gorestfull/models"
	"gopkg.in/go-playground/validator.v9"
)

type users struct {
	rp interfaces.UserServices
}

func User(rps interfaces.UserServices) *users {
	return &users{rps}
}

var validate *validator.Validate

func (us *users) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	data, err := us.rp.GetAll()

	if err != nil {
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// helpers.Response(w, &data, 200, false)
	data.Send(w)
}

func (us *users) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access.Control.Allow.Origin", "*")

	validate := validator.New()

	var body models.User

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	err = validate.Struct(body)

	if err != nil {
		helpers.Response(w, err.(validator.ValidationErrors).Error(), 400, true)
	}

	data := models.CreateUser()
	data.Id = body.Id
	data.Name = body.Name
	data.Username = body.Username
	data.Email = body.Email
	data.Role = body.Role

	hash, err := helpers.HashPassword(body.Password)

	if err != nil {
		fmt.Println("Hashing Password Failed")
	}

	data.Password = hash

	res, err := us.rp.Add(data)

	if err != nil {
		helpers.Response(w, err.Error(), 500, true)
		return
	}

	// helpers.Response(w, &data, 201, false)
	res.Send(w)
}

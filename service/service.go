package service

import (
	"simple_server/database"
	custom_errors "simple_server/errors"
	"simple_server/model"
)

func SignUp(user model.User) error {
	ok, _ := database.CheckLogin(user.Login)

	if ok {
		return custom_errors.ErrUserExists
	}

	err := database.SignUp(user)
	return err
}

func GetUsers() ([]model.User, error) {
	users, err := database.GetUsers()
	return users, err
}

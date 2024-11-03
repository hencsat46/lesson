package service

import (
	"errors"
	"simple_server/database"
	"simple_server/model"
)

func SignUp(user model.User) error {
	ok, _ := database.CheckLogin(user.Login)

	if ok {
		return errors.New("user exists")
	}

	err := database.SignUp(user)
	return err
}

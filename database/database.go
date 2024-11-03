package database

import "simple_server/model"

var Database map[string]string

// Заносим пользователя в базу данных
func SignUp(user model.User) error {
	Database[user.Login] = user.Password
	return nil
}

// Есть логин или нет
func CheckLogin(login string) (bool, error) {
	_, ok := Database[login]

	if ok {
		return true, nil
	}
	return false, nil
}

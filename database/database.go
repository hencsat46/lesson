package database

import (
	"fmt"
	"simple_server/model"
)

var Database map[string]string

func InitDb() {
	Database = make(map[string]string)
}

// Заносим пользователя в базу данных
func SignUp(user model.User) error {
	Database[user.Login] = user.Password
	fmt.Println(Database)
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

func GetUsers() ([]model.User, error) {
	usersArray := make([]model.User, 0)

	for login, password := range Database {
		temp := model.User{Login: login, Password: password}
		usersArray = append(usersArray, temp)
	}

	return usersArray, nil
}

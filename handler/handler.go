package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	custom_errors "simple_server/errors"
	"simple_server/model"
	"simple_server/service"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	request := model.User{}

	fmt.Println(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, 422, "Братанчик, ну ты и хуйню написал")
		return
	}

	if request.Login == "" || request.Password == "" {
		writeResponse(w, 422, "Json не соответствует схеме")
		return
	}

	if err := service.SignUp(request); err != nil {
		if errors.Is(err, custom_errors.ErrUserExists) {
			writeResponse(w, 422, "Пользователь уже существует")
			return
		}

		writeResponse(w, 500, "Ошибка сервера")
		return
	}

	writeResponse(w, 200, "Пользователь сохранён")
}

func SignIn(w http.ResponseWriter, r *http.Request) {

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetUsers()

	if err != nil {
		writeResponse(w, 500, "Internal server error")
		return
	}

	writeResponse(w, 200, users)
}

func writeResponse(w http.ResponseWriter, status int, content any) {
	response := model.Response{
		Status:  status,
		Content: content,
	}
	jsonString, _ := json.Marshal(response)
	w.Write(jsonString)
}

package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

	fmt.Println(request)

	err := service.SignUp(request)

	if errors.Is(err, errors.New("user exists")) {

	}

}

func SignIn(w http.ResponseWriter, r *http.Request) {

}

func writeResponse(w http.ResponseWriter, status int, content string) {
	response := model.Response{
		Status:  status,
		Content: content,
	}
	jsonString, _ := json.Marshal(response)
	w.Write(jsonString)
}

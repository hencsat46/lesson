package main

import (
	"net/http"
	"simple_server/database"
	"simple_server/handler"
)

func main() {
	database.InitDb()
	http.HandleFunc("POST /signup", handler.SignUp)
	http.HandleFunc("POST /signin", handler.SignIn)
	http.HandleFunc("GET /get_users", handler.GetUsers)
	http.ListenAndServe(":3000", nil)
}

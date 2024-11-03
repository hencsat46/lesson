package main

import (
	"net/http"
	"simple_server/handler"
)

func main() {
	http.HandleFunc("POST /signup", handler.SignUp)
	http.HandleFunc("POST /signin", handler.SignIn)
	http.ListenAndServe(":3000", nil)
}

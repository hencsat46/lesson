package model

type User struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
}

type Response struct {
	Status  int    `json:"Status"`
	Content string `json:"Content"`
}

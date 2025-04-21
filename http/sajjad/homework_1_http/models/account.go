package models

type Account struct {
	Attributes AccountAttributes `json:"attributes"`
}

type AccountAttributes struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

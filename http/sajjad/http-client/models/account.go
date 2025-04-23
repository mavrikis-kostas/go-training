package models

type Account struct {
	Attributes AccountAttributes `json:"attributes"`
}

type AccountAttributes struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

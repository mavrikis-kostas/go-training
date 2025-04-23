package useraccount

// Account represents a user account with attributes
type Account struct {
	Attributes AccountAttributes `json:"attributes"`
}

// AccountAttributes contains the account's attributes
type AccountAttributes struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

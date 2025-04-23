package user

// User represents a user with attributes
type User struct {
	Attributes Attributes `json:"attributes"`
}

// Attributes contains the user's attributes
type Attributes struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	AccountIDs []int  `json:"account_ids"`
}

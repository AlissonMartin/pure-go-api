package main


type CreateAccountRequest struct {
	Name string `json:"name"`
}

type Account struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Balance int64 `json:"balance"`
}

func NewAccount(Name string) *Account{
	return &Account{
		Name: Name,
	}
}
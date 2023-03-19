package models

type User struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"-"`
}

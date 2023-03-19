package models

type RegisterUser struct {
	Login    string `json:"login" validate:"required,alphanum,gte=3"` //не короче 3 символов, из цифр и английского алфавита
	Password string `json:"password" validate:"required,gte=6"`       //не короче 6 символов
}

type LoginUser struct {
	Login    string `json:"login" validate:"required,alphanum,gte=3"` //не короче 3 символов, из цифр и английского алфавита
	Password string `json:"password" validate:"required,gte=6"`       //не короче 6 символов
}

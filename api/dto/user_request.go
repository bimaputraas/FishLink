package dto

type ReqUserRegister struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string	`json:"password"`
	Address      string	`json:"address"`
	Phone        string	`json:"phone"`
}
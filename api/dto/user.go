package dto

type ReqBodyRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type ReqBodyLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserEmailVerification struct {
	Email            string `json:"email"`
	UserId           uint   `json:"user_id"`
	VerificationCode string `json:"verification_code"`
}

type TopUpReqBody struct {
	Amount int64 `json:"amount"`
}
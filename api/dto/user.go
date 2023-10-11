package dto

type ReqBodyRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type ReqBodyLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEmailVerification struct {
	Email            string `json:"email"`
	UserId           uint   `json:"user_id"`
	VerificationCode string `json:"verification_code"`
}

type ResBodyGetInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

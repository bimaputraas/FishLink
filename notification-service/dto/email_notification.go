package dto

type UserEmailVerification struct {
	Email            string `json:"email"`
	UserId           uint   `json:"user_id"`
	VerificationCode string `json:"verification_code"`
}

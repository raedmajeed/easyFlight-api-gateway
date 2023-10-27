package DTO

type OTP struct {
	Otp   int    `json:"OTP" validator:"required"`
	Email string `json:"email" validator:"required,email"`
}

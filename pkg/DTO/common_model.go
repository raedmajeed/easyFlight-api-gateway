package DTO

type OTP struct {
	Otp   int    `json:"OTP" validator:"required,numeric,len=6"`
	Email string `json:"email" validator:"required,email"`
}

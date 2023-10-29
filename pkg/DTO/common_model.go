package DTO

type OTP struct {
	Otp   int    `json:"OTP" validator:"required,numeric,len=6"`
	Email string `json:"email" validator:"required,email"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"password,min=8"`
}

type ConfirmPasswordRequest struct {
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=password"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email,required"`
}

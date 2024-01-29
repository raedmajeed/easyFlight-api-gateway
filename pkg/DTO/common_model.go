package DTO

type OTP struct {
	Otp   int    `json:"OTP" validator:"required,numeric"`
	Email string `json:"email" validator:"required,email"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,emailcst"`
	Password string `json:"password" validate:"required"`
}

type ConfirmPasswordRequest struct {
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email,required"`
}

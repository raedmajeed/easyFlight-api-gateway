package DTO

type UserData struct {
	Email    string `validate:"required,emailcst"`
	Phone    string `validate:"required,phone"`
	Password string `validate:"required,min=8"`
	Name     string `validate:"required,alphaspace"`
}

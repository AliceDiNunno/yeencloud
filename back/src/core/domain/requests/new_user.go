package requests

type NewUser struct {
	Email    string `json:"email" validate:"required,email,unique_email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

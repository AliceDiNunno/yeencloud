package requests

type NewSession struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	OTPCode  string `json:"otp_code"`

	//Field to be filled by server
	IP string `json:"-"`
}

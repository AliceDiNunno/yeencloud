package requests

type NewSession struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	OTPCode  string `json:"otpCode"`

	// Field to be filled by server
	IP string `json:"-"`
}

// Secure : remove sensitive data from the request.
func (n NewSession) Secure() NewSession {
	n.Password = ""
	// OTP code should be safe as it is time sensitive.
	return n
}

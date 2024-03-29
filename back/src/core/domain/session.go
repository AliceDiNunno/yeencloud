package domain

type Session struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expireAt"`
	IP       string `json:"ip"`
	UserID   UserID `json:"userId"`
}

type NewSession struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	OTPCode  string `json:"otpCode"`

	// Origin of the request to be filled internally (can, and will probably be an IP)
	Origin string `json:"-"`
}

// Secure : remove sensitive data from the request.
func (n NewSession) Secure() NewSession {
	n.Password = ""
	// OTP code should be safe as it is time sensitive.
	return n
}

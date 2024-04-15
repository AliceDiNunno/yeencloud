package domain

// MARK: - Objects

type TokenType string

type TokenID string

type Token struct {
	ID TokenID `json:"id"`

	User  User   `json:"user"`
	Token string `json:"token"`

	CreatedAt int64 `json:"createdAt"`
	ExpireAt  int64 `json:"expireAt"`

	Type TokenType `json:"type"`
}

var (
	TokenTypeVerifyEmail   = TokenType("verify_email")
	TokenTypeResetPassword = TokenType("reset_password")
)

// MARK: - Requests

type RequestNewPassword struct {
	Email string `json:"email" validate:"required,email"`
}

type RecoverPassword struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token" validate:"required"`
}

type ValidateMail struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token" validate:"required"`
}

// MARK: - Functions

func (t TokenType) String() string {
	return string(t)
}

func (t TokenID) String() string {
	return string(t)
}

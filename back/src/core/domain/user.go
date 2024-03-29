package domain

type UserID string

func InvalidUserID() UserID {
	return "00000000-0000-0000-0000-000000000000"
}

func (id UserID) String() string {
	return string(id)
}

// A user represents only the user's authentication data and maybe the email used for communication (up to further changes)
// The rest of the user's data will be found in the profile

type User struct {
	ID       UserID `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // Password (even if it is hashed) should never be exposed
}

type NewUser struct {
	Email    string `json:"email" validate:"required,email,unique_email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

func (u NewUser) Secure() NewUser {
	u.Password = ""
	return u
}

type UpdateUser struct {
	Email    string `json:"email"  validate:"email,unique_email"`
	Password string `json:"password" validate:"password"`
}

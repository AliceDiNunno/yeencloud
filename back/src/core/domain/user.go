package domain

type UserID string

func (id UserID) String() string {
	return string(id)
}

type User struct {
	ID       UserID
	Email    string
	Password string `json:"-"` // Password (even if it is hashed) should never be exposed
}

type Profile struct {
	UserID   UserID
	Name     string
	Language string
}

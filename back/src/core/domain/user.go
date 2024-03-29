package domain

type User struct {
	ID       string
	Email    string
	Password string `json:"-"` // Password (even if it is hashed) should never be exposed
}

type Profile struct {
	UserID   string
	Name     string
	Language string
}

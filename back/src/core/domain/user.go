package domain

type User struct {
	CloudObject

	Email string
}

type Profile struct {
	CloudObject

	UserID string
	Name   string
}

package domain

type UserID string
type ProfileID string

func (id UserID) String() string {
	return string(id)
}

func (id ProfileID) String() string {
	return string(id)
}

// A user represents only the user's authentication data and maybe the email used for communication (up to further changes)
// The rest of the user's data will be found in the profile

type User struct {
	ID       UserID `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // Password (even if it is hashed) should never be exposed
}

// A profile represents the user's profile and everything that is not related to authentication
// So this is what we're referencing when we want to get organizations, services, settings, etc...
// As in the future the user will be moved to an authentication service, we want to keep them isolated

type Profile struct {
	ID       ProfileID `json:"profileId"`
	UserID   UserID    `json:"userId"`
	Name     string    `json:"name"`
	Language string    `json:"language"`
}

package domain

type ProfileID string

func (id ProfileID) String() string {
	return string(id)
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

type UpdateProfile struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}

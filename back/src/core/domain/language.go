package domain

type Language struct {
	Identifier  string `json:"id"`
	Flag        string `json:"flag"`
	DisplayName string `json:"displayName"`
}

// Translation keys.
const (
	DefaultOrganizationDescription    = "DefaultOrganizationDescription"
	DefaultOrganizationDescriptionKey = "UserFullName"
)

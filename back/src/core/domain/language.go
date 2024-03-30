package domain

type Language struct {
	Tag         string `json:"tag"`
	Flag        string `json:"flag"`
	DisplayName string `json:"displayName"`
}

// Translation keys.
const (
	DefaultOrganizationDescription    = "DefaultOrganizationDescription"
	DefaultOrganizationDescriptionKey = "UserFullName"
)

package domain

type Language struct {
	Identifier  string
	Flag        string
	DisplayName string
}

// Translation keys.
const (
	DefaultOrganizationDescription    = "DefaultOrganizationDescription"
	DefaultOrganizationDescriptionKey = "UserFullName"
)

package domain

type OrganizationID string

func (id OrganizationID) String() string {
	return string(id)
}

type Organization struct {
	ID          OrganizationID `json:"id"`
	Slug        string         `json:"slug"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
}

type OrganizationRole string

const (
	OrganizationRoleOwner OrganizationRole = "OWNER"
	OrganizationRoleAdmin OrganizationRole = "ADMIN"
	OrganizationRoleUser  OrganizationRole = "USER"
)

func (id OrganizationRole) String() string {
	return string(id)
}

type OrganizationMember struct {
	Profile      Profile          `json:"profile"`
	Organization Organization     `json:"organization"`
	Role         OrganizationRole `json:"role"`
}

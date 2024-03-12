package domain

type OrganizationID string

func (id OrganizationID) String() string {
	return string(id)
}

type Organization struct {
	ID          OrganizationID
	Slug        string
	Name        string
	Description string
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
	User         User
	Organization Organization
	Role         OrganizationRole
}

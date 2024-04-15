package domain

// MARK: - Objects

type OrganizationMember struct {
	Profile      Profile      `json:"profile"`
	Organization Organization `json:"organization"`
	Role         string       `json:"role"`
}

// MARK: - Permissions

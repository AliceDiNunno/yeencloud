package domain

// MARK: - Objects

var roleSeparator = ":"

type RoleScope struct {
	Parent *RoleScope

	Identifier  string
	DisplayName string
}

// MARK: - Functions

func (r RoleScope) String() string {
	if r.Parent != nil {
		return r.Parent.String() + roleSeparator + r.Identifier
	}
	return r.Identifier
}

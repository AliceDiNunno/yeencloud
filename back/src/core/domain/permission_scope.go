package domain

// MARK: - Objects

type PermissionScope struct {
	Parent *PermissionScope

	Name string
}

// MARK: - Functions

func (d PermissionScope) Identifier() string {
	if d.Parent != nil {
		return d.Parent.Identifier() + ":" + d.Name
	}
	return d.Name
}

package domain

import "net/http"

// MARK: - Objects

type Permission struct {
	Parent PermissionScope

	Name string
}

// MARK: - Translatable Args
var (
	TranslatableArgumentPermissionIdentifier = TranslatableArgument{Key: "PermissionIdentifier"}
)

// MARK: - Translatable
var (
	TranslatablePermissionRequired = Translatable{Key: "PermissionRequired", Arguments: []TranslatableArgument{TranslatableArgumentPermissionIdentifier}}
)

// MARK: - Errors
var (
	errorPermissionRequired = ErrorDescription{HttpCode: http.StatusForbidden, Code: TranslatablePermissionRequired}
)

// MARK: - Functions

func (p Permission) Identifier() string {
	return p.Parent.Identifier() + ":" + p.Name
}

func PermissionByIdentifier(id string) *Permission {
	for _, permission := range AllPermissions() {
		if permission.Identifier() == id {
			return &permission
		}
	}
	return nil
}

func ErrorPermissionRequired(permission Permission) *ErrorDescription {
	errorPermissionRequired.Arguments = TranslatableArgumentMap{TranslatableArgumentPermissionIdentifier: permission.Identifier()}
	return &errorPermissionRequired
}

// MARK: - Permissions
// When adding a new permission, don't forget to declare it to the AllPermissions function.
// Not doing so may result in unexpected behavior. (However this will be tested in the future #TODO )

func AllPermissions() []Permission {
	var permissions []Permission

	permissions = append(permissions, OrganizationPermissions...)

	return permissions
}

package domain

import (
	"fmt"
)

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

// MARK: - Resource Not Found

type PermissionRequiredError struct {
	Permission string
}

func (e *PermissionRequiredError) Error() string {
	return fmt.Sprintf("usecases: permission required: '%v' to perform this operation", e.Permission)
}

func (e *PermissionRequiredError) RawKey() Translatable {
	return TranslatablePermissionRequired
}

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

func ErrorPermissionRequired(permission Permission) error {
	err := PermissionRequiredError{Permission: permission.Identifier()}
	return &err
}

// MARK: - Permissions
// When adding a new permission, don't forget to declare it to the AllPermissions function.
// Not doing so may result in unexpected behavior. (However this will be tested in the future #TODO )

func AllPermissions() []Permission {
	var permissions []Permission

	permissions = append(permissions, OrganizationPermissions...)

	return permissions
}

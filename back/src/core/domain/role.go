package domain

// MARK: - Objects

type Role struct {
	Scope       RoleScope
	Name        string
	DisplayName Translatable

	Inherit     []Role
	Permissions []Permission
}

// MARK: - Functions

func (r Role) PermissionsList() []Permission {
	permissions := r.Permissions
	for _, role := range r.Inherit {
		permissions = append(permissions, role.PermissionsList()...)
	}
	return permissions
}

func (r Role) HasPermission(permission Permission) bool {
	for _, p := range r.PermissionsList() {
		if p == permission {
			return true
		}
	}
	return false
}

func RoleByName(identifier string) Role {
	for _, role := range AllRoles() {
		if role.String() == identifier {
			return role
		}
	}
	return RoleInvalidRole
}

func (r Role) String() string {
	return r.Scope.String() + roleSeparator + r.Name
}

// MARK: - Translatable

var TranslatableRoleInvalidDisplayName = Translatable{Key: "RoleInvalidDisplayName"}

// MARK: - Roles

var RoleInvalidScope = RoleScope{
	Identifier: "invalid",
}

var RoleInvalidRole = Role{
	Scope:       RoleInvalidScope,
	Name:        "invalid",
	DisplayName: TranslatableRoleInvalidDisplayName,

	Permissions: []Permission{},
}

// When adding a new role, don't forget to declare it to the AllRoles function.
// Not doing so may result in unexpected behavior. (However this will be tested in the future #TODO )

func AllRoles() []Role {
	var roles []Role

	roles = append(roles, ProfileRoles...)
	roles = append(roles, OrganizationRoles...)

	return roles
}

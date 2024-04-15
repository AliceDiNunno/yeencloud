package domain

// MARK: - Objects

type OrganizationID string

type Organization struct {
	ID          OrganizationID `json:"id"`
	Slug        string         `json:"slug"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
}

// MARK: - Requests

type NewOrganization struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type UpdateOrganization struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// MARK: - Translatable
var (
	TranslatableOrganizationMemberRoleDisplayName = Translatable{Key: "OrganizationMemberRoleDisplayName"}
	TranslatableOrganizationAdminRoleDisplayName  = Translatable{Key: "OrganizationAdminRoleDisplayName"}
	TranslatableOrganizationOwnerRoleDisplayName  = Translatable{Key: "OrganizationOwnerRoleDisplayName"}
)

// MARK: - Functions

func (id OrganizationID) String() string {
	return string(id)
}

// MARK: - Permissions
var (
	PermissionScopeDomainOrganization     = PermissionScope{Name: "organization"}
	PermissionScopeDomainOrganizationMeta = PermissionScope{Parent: &PermissionScopeDomainOrganization, Name: "meta"}

	PermissionGlobalOrganizationCreation = Permission{Name: "create", Parent: PermissionScopeDomainOrganization}

	PermissionOrganizationMetaUpdate = Permission{Name: "update", Parent: PermissionScopeDomainOrganizationMeta}

	PermissionOrganizationRead   = Permission{Name: "read", Parent: PermissionScopeDomainOrganization}
	PermissionOrganizationDelete = Permission{Name: "delete", Parent: PermissionScopeDomainOrganization}
)

var OrganizationPermissions = []Permission{
	PermissionGlobalOrganizationCreation,

	PermissionOrganizationMetaUpdate,

	PermissionOrganizationRead,
	PermissionOrganizationDelete,
}

// MARK: - Roles

var RoleScopeOrganization = RoleScope{
	Identifier:  "organization",
	DisplayName: "Organization",
}

var RoleOrganizationMember = Role{
	Scope:       RoleScopeOrganization,
	Name:        "member",
	DisplayName: TranslatableOrganizationMemberRoleDisplayName,

	Permissions: []Permission{
		PermissionOrganizationRead,
	},
}

var RoleOrganizationAdmin = Role{
	Scope:       RoleScopeOrganization,
	Name:        "admin",
	DisplayName: TranslatableOrganizationAdminRoleDisplayName,

	Inherit: []Role{
		RoleOrganizationMember,
	},
	Permissions: []Permission{
		PermissionOrganizationMetaUpdate,
	},
}

var RoleOrganizationOwner = Role{
	Scope:       RoleScopeOrganization,
	Name:        "owner",
	DisplayName: TranslatableOrganizationOwnerRoleDisplayName,

	Inherit: []Role{
		RoleOrganizationAdmin,
	},
	Permissions: []Permission{
		PermissionOrganizationDelete,
	},
}

var OrganizationRoles = []Role{
	RoleOrganizationMember,
	RoleOrganizationAdmin,
	RoleOrganizationOwner,
}

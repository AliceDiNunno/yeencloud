package domain

// MARK: - Objects

type ProfileID string

// A profile represents the user's profile and everything that is not related to authentication
// So this is what we're referencing when we want to get organizations, services, settings, etc...
// As in the future the user will be moved to an authentication service, we want to keep them isolated

type Profile struct {
	ID       ProfileID `json:"profileId"`
	UserID   UserID    `json:"userId"`
	Name     string    `json:"name"`
	Language string    `json:"language"`
	Role     string    `json:"role"`
}

// MARK: - Requests

type UpdateProfile struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}

// MARK: - Translatable
var (
	TranslatableProfileRoleUnvalidatedDisplayName = Translatable{Key: "ProfileRoleUnvalidatedDisplayName"}
	TranslatableProfileRoleStandardDisplayName    = Translatable{Key: "ProfileRoleStandardDisplayName"}
	TranslatableProfileRoleStaffDisplayName       = Translatable{Key: "ProfileRoleStaffDisplayName"}
)

// MARK: - Logs
var (
	LogScopeProfile     = LogScope{Identifier: "profile"}
	LogFieldProfileID   = LogField{Scope: LogScopeProfile, Identifier: "id"}
	LogFieldProfileRole = LogField{Scope: LogScopeProfile, Identifier: "role"}
	LogFieldProfileMail = LogField{Scope: LogScopeProfile, Identifier: "mail"}
	LogFieldProfileName = LogField{Scope: LogScopeProfile, Identifier: "name"}
)

// MARK: - Functions

func (id ProfileID) String() string {
	return string(id)
}

// MARK: - Roles

var RoleScopeProfile = RoleScope{
	Identifier:  "profile",
	DisplayName: "Profile",
}

// RoleProfileUnvalidated is the role that is assigned if a profile is not validated yet.
var RoleProfileUnvalidated = Role{
	Scope: RoleScopeProfile,

	Name:        "unvalidated",
	DisplayName: TranslatableProfileRoleUnvalidatedDisplayName,

	Permissions: []Permission{},
}

var RoleProfileStandard = Role{
	Scope: RoleScopeProfile,

	Name:        "standard",
	DisplayName: TranslatableProfileRoleStandardDisplayName,
	Inherit: []Role{
		RoleProfileUnvalidated,
	},
	Permissions: []Permission{
		PermissionGlobalOrganizationCreation,
	},
}

// RoleProfileStaff is basically a "root" role that's used to define the permissions of an internal staff member
// it might also be used for internal scripts (tbd). Name makes it so it's not to be confused with "admin" or "owner" of other entities
// if you are in production and have this role, you should be able to do everything (with great power comes great responsibility).
var RoleProfileStaff = Role{
	Scope: RoleScopeProfile,

	Name:        "staff",
	DisplayName: TranslatableProfileRoleStaffDisplayName,
	Inherit: []Role{
		RoleProfileStandard,

		// those roles should not be inherited by any non-staff profile roles, they give owner power to everything
		RoleOrganizationOwner,
	},
	Permissions: []Permission{},
}

var ProfileRoles = []Role{
	RoleProfileUnvalidated,
	RoleProfileStandard,
	RoleProfileStaff,
}

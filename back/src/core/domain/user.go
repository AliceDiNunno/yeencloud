package domain

import "net/http"

// MARK: - Objects

type UserID string

// A user represents only the user's authentication data and maybe the email used for communication (up to further changes)
// The rest of the user's data will be found in the profile

type User struct {
	ID       UserID `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // Password (even if it is hashed) should never be exposed
}

// MARK: - Requests

type NewUser struct {
	Email    string `json:"email" validate:"required,email,unique_email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

type UpdateUser struct {
	Email    string `json:"email"  validate:"email,unique_email"`
	Password string `json:"password" validate:"password"`
}

// MARK: - Translatable
var (
	TranslatableUserNotFound                 = Translatable{Key: "UserNotFound"}
	TranslatableUnableToCreateUser           = Translatable{Key: "UnableToCreateUser"}
	TranslatableUnableToHashPassword         = Translatable{Key: "UnableToHashPassword"}
	TranslatableUnableToGetUserOrganizations = Translatable{Key: "UnableToGetUserOrganizations"}

	TranslatableUserRoleLimitedDisplayName  = Translatable{Key: "UserRoleLimitedDisplayName"}
	TranslatableUserRoleStandardDisplayName = Translatable{Key: "UserRoleStandardDisplayName"}
	TranslatableUserRoleStaffDisplayName    = Translatable{Key: "UserRoleStaffDisplayName"}
)

// MARK: - Errors
var (
	ErrorUserNotFound         = ErrorDescription{HttpCode: http.StatusNotFound, Code: TranslatableUserNotFound}
	ErrorUnableToCreateUser   = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToCreateUser}
	ErrorUnableToHashPassword = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToHashPassword}
)

// MARK: - Logs
var (
	LogFieldUser   = LogScope{Identifier: "user"}
	LogFieldUserID = LogField{Scope: LogFieldUser, Identifier: "id"}
)

// MARK: - Functions

func InvalidUserID() UserID {
	return "00000000-0000-0000-0000-000000000000"
}

func (id UserID) String() string {
	return string(id)
}

func (u NewUser) Secure() NewUser {
	u.Password = ""
	return u
}

// MARK: - Roles

var RoleScopeUser = RoleScope{
	Identifier:  "user",
	DisplayName: "User",
}

// RoleUserLimited is the role that is assigned if a user is not validated yet.
var RoleUserLimited = Role{
	Scope: RoleScopeUser,

	Name:        "limited",
	DisplayName: TranslatableUserRoleLimitedDisplayName,

	Permissions: []Permission{},
}

var RoleUserStandard = Role{
	Scope: RoleScopeUser,

	Name:        "standard",
	DisplayName: TranslatableUserRoleStandardDisplayName,
	Inherit: []Role{
		RoleUserLimited,
	},
	Permissions: []Permission{
		PermissionGlobalOrganizationCreation,
	},
}

// RoleUserStaff is basically a "root" role that's used to define the permissions of an internal staff member
// it might also be used for internal scripts (tbd). Name makes it so it's not to be confused with "admin" or "owner" of other entities
// if you are in production and have this role, you should be able to do everything (with great power comes great responsibility).
var RoleUserStaff = Role{
	Scope: RoleScopeUser,

	Name:        "staff",
	DisplayName: TranslatableUserRoleStaffDisplayName,
	Inherit: []Role{
		RoleUserStandard,

		// those roles should not be inherited by any non-staff user roles, they give owner power to everything
		RoleOrganizationOwner,
	},
	Permissions: []Permission{},
}

var UserRoles = []Role{
	RoleUserLimited,
	RoleUserStandard,
	RoleUserStaff,
}

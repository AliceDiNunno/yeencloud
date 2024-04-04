package domain

import "net/http"

type ErrorDescription struct {
	HttpCode int          `json:"-"`
	Code     Translatable `json:"code"`
}

var (
	TranslatableUserNotFound       = Translatable{Key: "UserNotFound"}
	TranslatableUnableToCreateUser = Translatable{Key: "UnableToCreateUser"}

	TranslatableProfileNotFound       = Translatable{Key: "ProfileNotFound"}
	TranslatableUnableToCreateProfile = Translatable{Key: "UnableToCreateProfile"}

	TranslatableUnableToHashPassword         = Translatable{Key: "UnableToHashPassword"}
	TranslatableUnableToGetUserOrganizations = Translatable{Key: "UnableToGetUserOrganizations"}
	TranslatableSessionNotFound              = Translatable{Key: "SessionNotFound"}

	TranslatableOrganizationNotFound       = Translatable{Key: "OrganizationNotFound"}
	TranslatableUnableToCreateOrganization = Translatable{Key: "UnableToCreateOrganization"}
	TranslatableUnableToUpdateOrganization = Translatable{Key: "UnableToUpdateOrganization"}
	TranslatableUnableToDeleteOrganization = Translatable{Key: "UnableToDeleteOrganization"}

	TranslatableUnableToGetOrganizationMembers = Translatable{Key: "UnableToGetOrganizationMembers"}
	TranslatableUnableToLinkUserOrganization   = Translatable{Key: "UnableToLinkUserOrganization"}
)

var (
	ErrorUserNotFound       = ErrorDescription{HttpCode: http.StatusNotFound, Code: TranslatableUserNotFound}
	ErrorUnableToCreateUser = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToCreateUser}

	ErrorProfileNotFound       = ErrorDescription{HttpCode: http.StatusNotFound, Code: TranslatableProfileNotFound}
	ErrorUnableToCreateProfile = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToCreateProfile}

	ErrorUnableToHashPassword         = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToHashPassword}
	ErrorUnableToGetUserOrganizations = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToGetUserOrganizations}
	ErrorSessionNotFound              = ErrorDescription{HttpCode: http.StatusUnauthorized, Code: TranslatableSessionNotFound}

	ErrorOrganizationNotFound       = ErrorDescription{HttpCode: http.StatusNotFound, Code: TranslatableOrganizationNotFound}
	ErrorUnableToCreateOrganization = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToCreateOrganization}
	ErrorUnableToUpdateOrganization = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToUpdateOrganization}
	ErrorUnableToDeleteOrganization = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToDeleteOrganization}

	ErrorUnableToGetOrganizationMembers = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToGetOrganizationMembers}
	ErrorUnableToLinkUserOrganization   = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToLinkUserOrganization}
)

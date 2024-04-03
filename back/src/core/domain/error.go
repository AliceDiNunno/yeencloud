package domain

import "net/http"

type ErrorDescription struct {
	HttpCode int          `json:"-"`
	Code     Translatable `json:"code"`
}

var (
	TranslatableUserNotFound = Translatable{Key: "UserNotFound"}

	TranslatableProfileNotFound       = Translatable{Key: "ProfileNotFound"}
	TranslatableUnableToCreateProfile = Translatable{Key: "UnableToCreateProfile"}

	TranslatableUnableToHashPassword           = Translatable{Key: "UnableToHashPassword"}
	TranslatableUnableToGetUserOrganizations   = Translatable{Key: "UnableToGetUserOrganizations"}
	TranslatableUnableToGetOrganizationMembers = Translatable{Key: "UnableToGetOrganizationMembers"}
	TranslatableUserAlreadyExists              = Translatable{Key: "UserAlreadyExists"}
	TranslatableSessionNotFound                = Translatable{Key: "SessionNotFound"}

	TranslatableOrganizationNotFound       = Translatable{Key: "OrganizationNotFound"}
	TranslatableUnableToUpdateOrganization = Translatable{Key: "UnableToUpdateOrganization"}
	TranslatableUnableToDeleteOrganization = Translatable{Key: "UnableToDeleteOrganization"}
)

var (
	ErrorUserNotFound = ErrorDescription{HttpCode: http.StatusNotFound, Code: TranslatableUserNotFound}

	ErrorProfileNotFound       = ErrorDescription{HttpCode: http.StatusNotFound, Code: TranslatableProfileNotFound}
	ErrorUnableToCreateProfile = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToCreateProfile}

	ErrorUnableToHashPassword           = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToHashPassword}
	ErrorUnableToGetUserOrganizations   = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToGetUserOrganizations}
	ErrorUnableToGetOrganizationMembers = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToGetOrganizationMembers}
	ErrorUserAlreadyExists              = ErrorDescription{HttpCode: http.StatusConflict, Code: TranslatableUserAlreadyExists}
	ErrorSessionNotFound                = ErrorDescription{HttpCode: http.StatusUnauthorized, Code: TranslatableSessionNotFound}

	ErrorOrganizationNotFound       = ErrorDescription{HttpCode: http.StatusNotFound, Code: TranslatableOrganizationNotFound}
	ErrorUnableToUpdateOrganization = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToUpdateOrganization}
	ErrorUnableToDeleteOrganization = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToDeleteOrganization}
)

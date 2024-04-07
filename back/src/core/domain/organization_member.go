package domain

import "net/http"

// MARK: - Objects

type OrganizationMember struct {
	Profile      Profile      `json:"profile"`
	Organization Organization `json:"organization"`
	Role         string       `json:"role"`
}

// MARK: - Translatable

var (
	TranslatableUnableToGetOrganizationMembers    = Translatable{Key: "UnableToGetOrganizationMembers"}
	TranslatableUnableToLinkUserOrganization      = Translatable{Key: "UnableToLinkUserOrganization"}
	TranslatableUnableToRemoveOrganizationMembers = Translatable{Key: "UnableToRemoveOrganizationMembers"}
	TranslatableUnableToGetOrganizationRole       = Translatable{Key: "UnableToGetOrganizationRole"}
	TranslatableNotAuthorizedToModifyOrganization = Translatable{Key: "NotAuthorizedToModifyOrganization"}
)

// MARK: - Errors

var (
	ErrorNotAuthorizedToModifyOrganization = ErrorDescription{HttpCode: http.StatusForbidden, Code: TranslatableNotAuthorizedToModifyOrganization}
	ErrorUnableToGetUserOrganizations      = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToGetUserOrganizations}
	ErrorUnableToGetOrganizationMembers    = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToGetOrganizationMembers}
	ErrorUnableToLinkUserOrganization      = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToLinkUserOrganization}
	ErrorUnableToRemoveOrganizationMembers = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToRemoveOrganizationMembers}
	ErrorUnableToGetOrganizationRole       = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToGetOrganizationRole}
)

// MARK: - Permissions

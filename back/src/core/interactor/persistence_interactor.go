package interactor

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)

	CountUsers() int64
	FindUserByID(userID domain.UserID) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
}

type ProfileRepository interface {
	CreateProfile(profile domain.Profile) (domain.Profile, error)

	FindProfileByUserID(userID domain.UserID) (domain.Profile, error)
}

type OrganizationRepository interface {
	CreateOrganization(organization domain.Organization) (domain.Organization, error)

	UpdateOrganization(organization domain.OrganizationID, update domain.UpdateOrganization) (domain.Organization, error)

	DeleteOrganizationByID(id domain.OrganizationID) error
}

type OrganizationProfileRepository interface {
	LinkProfileToOrganization(profileID domain.ProfileID, organizationID domain.OrganizationID, role domain.OrganizationRole) error

	GetProfileOrganizationsByProfileID(profileID domain.ProfileID) ([]domain.OrganizationMember, error)
	GetOrganizationMembers(organizationID domain.OrganizationID) ([]domain.OrganizationMember, error)
}

type ServiceRepository interface {
}

type SessionRepository interface {
	CreateSession(session domain.Session) (domain.Session, error)

	FindSessionByToken(token string) (domain.Session, error)
}

type Persistence struct {
	// Main models
	User         UserRepository
	Service      ServiceRepository
	Session      SessionRepository
	Profile      ProfileRepository
	Organization OrganizationRepository

	// Linking models
	OrganizationProfile OrganizationProfileRepository
}

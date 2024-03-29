package persistence

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

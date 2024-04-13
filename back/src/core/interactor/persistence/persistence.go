package persistence

type Persistence interface {
	// Main models
	UserRepository
	ServiceRepository
	SessionRepository
	ProfileRepository
	OrganizationRepository
	TokenRepository

	// Linking models
	OrganizationProfileRepository

	// Transaction
	Transaction
}

package persistence

type Persistence interface {
	// Main models
	UserRepository
	ServiceRepository
	SessionRepository
	ProfileRepository
	OrganizationRepository

	// Linking models
	OrganizationProfileRepository

	// Transaction
	Transaction
}

package usecases

type Usecases interface {
	UserUsecases
	ProfileUsecases
	SessionUsecases
	OrganizationUsecases
	TransactionRequest
}

package persistence

type Transaction interface {
	Begin() Persistence // returns a new persistence instance with a transaction

	Commit() error
	Rollback() error
}

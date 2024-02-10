package postgres

import (
	"back/src/core/domain"
	"github.com/google/uuid"
)

func (db *Database) CountUsers() int64 {
	return 0
}

func (db *Database) FindUserByID(id uuid.UUID) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

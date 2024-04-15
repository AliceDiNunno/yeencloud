package postgres

import (
	"fmt"

	"github.com/AliceDiNunno/yeencloud/internal/core/domain/config"
	"github.com/AliceDiNunno/yeencloud/internal/core/interactor"
	"github.com/AliceDiNunno/yeencloud/internal/core/interactor/persistence"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	engine *gorm.DB
}

func (db *Database) Begin() persistence.Persistence {
	newTransaction := db.engine.Begin()

	if newTransaction.Error != nil {
		panic(newTransaction.Error)
	}

	return &Database{
		engine: newTransaction,
	}
}

func (db *Database) Commit() error {
	return db.engine.Commit().Error
}

func (db *Database) Rollback() error {
	return db.engine.Rollback().Error
}

func StartGormDatabase(log interactor.Logger, config config.DatabaseConfig) (*Database, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)

	db, err := gorm.Open(pg.Open(psqlInfo), &gorm.Config{
		Logger: newGormLogger(log),
	})

	if err != nil {
		return nil, sqlerr(err)
	}

	return &Database{
		engine: db,
	}, nil
}

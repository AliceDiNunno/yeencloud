package postgres

import (
	"fmt"

	"github.com/AliceDiNunno/yeencloud/src/core/domain/config"
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	engine *gorm.DB
}

func StartGormDatabase(log interactor.Logger, config config.DatabaseConfig) (*Database, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)

	db, err := gorm.Open(pg.Open(psqlInfo), &gorm.Config{
		Logger: newGormLogger(log),
	})

	if err != nil {
		return nil, err
	}

	return &Database{
		engine: db,
	}, nil
}

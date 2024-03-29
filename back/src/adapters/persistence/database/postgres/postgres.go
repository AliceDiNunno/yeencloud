package postgres

import (
	"back/src/core/domain/config"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	log "github.com/rs/zerolog/log"
	"github.com/wei840222/gorm-zerolog"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	engine *gorm.DB
}

func StartGormDatabase(config config.DatabaseConfig) (*Database, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)

	db, err := gorm.Open(pg.Open(psqlInfo), &gorm.Config{
		Logger: gorm_zerolog.NewWithLogger(log.Logger),
	})

	if err != nil {
		return nil, err
	}

	db.Logger.LogMode(logger.Info)

	return &Database{
		engine: db,
	}, nil
}

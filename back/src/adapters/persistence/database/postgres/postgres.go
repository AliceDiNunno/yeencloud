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
	engine      *gorm.DB
	serviceName string
}

func StartGormDatabase(config config.DatabaseConfig, serviceName string) *Database {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)

	db, err := gorm.Open(pg.Open(psqlInfo), &gorm.Config{
		Logger: gorm_zerolog.NewWithLogger(log.Logger),
	})

	if err != nil {
		panic(err)
	}
	db.Logger.LogMode(logger.Info)

	return &Database{
		engine:      db,
		serviceName: serviceName,
	}
}

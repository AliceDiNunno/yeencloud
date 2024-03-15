package postgres

import "github.com/rs/zerolog/log"

func (db *Database) Migrate() error {
	log.Info().Msg("Migrating database")

	err := db.engine.Debug().AutoMigrate(User{}, Organization{}, Profile{}, Session{})
	if err != nil {
		log.Err(err).Msg("Error migrating models")
		return err
	}

	err = db.engine.Debug().AutoMigrate(&OrganizationProfile{})
	if err != nil {
		log.Err(err).Msg("Error migrating linking tables")
		return err
	}

	log.Info().Msg("Database migrated")
	return nil
}

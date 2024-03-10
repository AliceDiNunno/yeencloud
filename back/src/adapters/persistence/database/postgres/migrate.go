package postgres

import "github.com/rs/zerolog/log"

func (db *Database) Migrate() {
	log.Info().Msg("Migrating database")

	err := db.engine.Debug().AutoMigrate(Settings{}, User{}, Organization{}, Profile{}, Session{})
	if err != nil {
		log.Err(err).Msg("Error migrating database")
		return
	}
}

package postgres

func (db *Database) Migrate() error {
	err := db.engine.Debug().AutoMigrate(User{}, Organization{}, Profile{}, Session{})
	if err != nil {
		//	log.Err(err).Msg("Error migrating models")
		return err
	}

	err = db.engine.Debug().AutoMigrate(&OrganizationProfile{})
	if err != nil {
		/*db.logger.Info(context.TODO(), "Error migrating linking tables", map[string]interface{}{
			"err": err,
		})*/

		//log.Err(err).Msg("Error migrating linking tables")
		return err
	}
	return nil
}

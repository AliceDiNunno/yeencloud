package postgres

func (db *Database) Migrate() error {
	err := db.engine.Debug().AutoMigrate(User{}, Organization{}, Profile{}, Session{})
	if err != nil {
		return err
	}

	err = db.engine.Debug().AutoMigrate(&OrganizationProfile{})
	if err != nil {
		return err
	}
	return nil
}

package postgres

func (db *Database) Migrate() error {
	err := db.engine.Debug().AutoMigrate(User{}, Organization{}, Profile{}, Session{}, Token{})
	if err != nil {
		return sqlerr(err)
	}

	err = db.engine.Debug().AutoMigrate(&OrganizationProfile{})
	if err != nil {
		return sqlerr(err)
	}
	return nil
}

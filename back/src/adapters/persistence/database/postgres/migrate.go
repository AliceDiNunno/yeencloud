package postgres

func (db *Database) Migrate() {
	err := db.engine.AutoMigrate(Settings{})
	if err != nil {
		return
	}
}

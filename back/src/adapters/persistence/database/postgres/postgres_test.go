package postgres

import (
	testConf "back/src/core/config/testConfig"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DatabaseTestSuite struct {
	suite.Suite
	database *Database
}

func (suite *DatabaseTestSuite) SetupTest() {
	config := testConf.NewConfig()
	config.GetDatabaseConfig()

	database, err := StartGormDatabase(config.GetDatabaseConfig())

	if err == nil {
		database.Migrate()
		suite.database = database
	}
}

func TestDatabaseSuiteIntegration(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}

package postgres

import (
	testConf "back/src/core/config/testConfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type DatabaseIntegrationTestSuite struct {
	suite.Suite
	database *Database
}

type DatabaseDomainBridgeTestSuite struct {
	suite.Suite
}

func globalTearDown(database *Database, t *testing.T) {
	// The tables are deleted in the following order to avoid foreign key constraints errors

	// Linking tables
	err := database.engine.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&OrganizationProfile{}).Error
	assert.NoError(t, err)

	// Models
	err = database.engine.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&Organization{}).Error
	assert.NoError(t, err)
	err = database.engine.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&Profile{}).Error
	assert.NoError(t, err)
	err = database.engine.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&Session{}).Error
	assert.NoError(t, err)
	err = database.engine.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&User{}).Error
	assert.NoError(t, err)
}

func (suite *DatabaseIntegrationTestSuite) TearDownTest() {
	globalTearDown(suite.database, suite.T())
}

func (suite *DatabaseIntegrationTestSuite) migrateTestDB() error {
	config := testConf.NewConfig()
	config.GetDatabaseConfig()

	database, err := StartGormDatabase(config.GetDatabaseConfig())
	if err != nil {
		return err
	}

	if err == nil {
		err = database.Migrate()
		suite.database = database
		return err
	}

	return nil
}

func TestDatabaseSuite(t *testing.T) {
	//Testing toDomain and fromDomain
	suite.Run(t, new(DatabaseDomainBridgeTestSuite))

	//Testing database (those are integration tests so they are not run in short mode)
	if !testing.Short() {
		integrationTests := new(DatabaseIntegrationTestSuite)
		err := integrationTests.migrateTestDB()
		assert.NoError(t, err)

		suite.Run(t, integrationTests)

		userModelIntegrationTestSuite := new(UserModelIntegrationTestSuite)
		userModelIntegrationTestSuite.database = integrationTests.database

		profileModelIntegrationTestSuite := new(ProfileModelIntegrationTestSuite)
		profileModelIntegrationTestSuite.database = integrationTests.database

		sessionModelIntegrationTestSuite := new(SessionModelIntegrationTestSuite)
		sessionModelIntegrationTestSuite.database = integrationTests.database

		organizationModelIntegrationTestSuite := new(OrganizationModelIntegrationTestSuite)
		organizationModelIntegrationTestSuite.database = integrationTests.database

		organizationProfileModelIntegrationTestSuite := new(OrganizationProfileModelIntegrationTestSuite)
		organizationProfileModelIntegrationTestSuite.database = integrationTests.database

		suite.Run(t, userModelIntegrationTestSuite)
		suite.Run(t, profileModelIntegrationTestSuite)
		suite.Run(t, sessionModelIntegrationTestSuite)
		suite.Run(t, organizationModelIntegrationTestSuite)

		suite.Run(t, organizationProfileModelIntegrationTestSuite)
	}
}

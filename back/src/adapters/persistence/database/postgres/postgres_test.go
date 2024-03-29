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

func (suite *DatabaseIntegrationTestSuite) TestFirst() {
	println("first")
}

func (suite *DatabaseIntegrationTestSuite) TestSecond() {
	println("second")
}

func (suite *DatabaseIntegrationTestSuite) TearDownTest() {
	// The tables are deleted in the following order to avoid foreign key constraints errors

	// Linking tables
	err := suite.database.engine.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&OrganizationProfile{}).Error
	suite.NoError(err)

	// Models
	err = suite.database.engine.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&Organization{}).Error
	suite.NoError(err)
	err = suite.database.engine.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&Profile{}).Error
	suite.NoError(err)
	err = suite.database.engine.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&Session{}).Error
	suite.NoError(err)
	err = suite.database.engine.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&User{}).Error
	suite.NoError(err)

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
	}
}

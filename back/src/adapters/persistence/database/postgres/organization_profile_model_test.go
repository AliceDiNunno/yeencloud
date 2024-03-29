package postgres

import "github.com/stretchr/testify/suite"

type OrganizationProfileModelIntegrationTestSuite struct {
	suite.Suite
	database *Database
}

func (suite *OrganizationProfileModelIntegrationTestSuite) SetupTest() {

}

func (suite *OrganizationProfileModelIntegrationTestSuite) TearDownTest() {
	globalTearDown(suite.database, suite.T())
}

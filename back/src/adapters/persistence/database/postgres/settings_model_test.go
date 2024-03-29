package postgres

import "testing"

func (suite *DatabaseTestSuite) TestStartGormDatabaseIntegration() {
	if testing.Short() {
		suite.T().Skip("skipping integration test")
	}

	suite.Assert().Equal("A", "B")
}

package postgres

import (
	"back/src/core/domain"
	"github.com/stretchr/testify/suite"
)

type ProfileModelIntegrationTestSuite struct {
	suite.Suite
	database *Database
}

var testProfile = domain.Profile{
	ID:       "b6940a9e-2e70-4c03-9b36-ee0642dd5ce1",
	UserID:   testUser.ID,
	Name:     "Jean-Michel Micheline",
	Language: "en-US",
}

func (suite *DatabaseDomainBridgeTestSuite) TestProfileToDomain() {
	//Given
	modelProfile := Profile{
		ID:       testProfile.ID.String(),
		UserID:   testProfile.UserID.String(),
		Name:     testProfile.Name,
		Language: testProfile.Language,
	}

	//When
	domainProfile := profileToDomain(modelProfile)

	//Then
	suite.Assert().Equal(modelProfile.ID, domainProfile.ID.String())
	suite.Assert().Equal(modelProfile.UserID, domainProfile.UserID.String())
	suite.Assert().Equal(modelProfile.Name, domainProfile.Name)
	suite.Assert().Equal(modelProfile.Language, domainProfile.Language)
}

func (suite *DatabaseDomainBridgeTestSuite) TestDomainToProfile() {
	//Given
	domainProfile := testProfile

	//When
	modelProfile := domainToProfile(domainProfile)

	//Then
	suite.Assert().Equal(testProfile.ID.String(), modelProfile.ID)
	suite.Assert().Equal(testProfile.UserID.String(), modelProfile.UserID)
	suite.Assert().Equal(testProfile.Name, modelProfile.Name)
	suite.Assert().Equal(testProfile.Language, modelProfile.Language)
}

func (suite *DatabaseDomainBridgeTestSuite) TestProfileDomainToModelToDomain() {
	//Given
	domainProfile := testProfile

	//When
	profile := domainToProfile(domainProfile)
	profileDomain := profileToDomain(profile)

	//Then
	suite.Assert().Equal(domainProfile, profileDomain)
}

func (suite *ProfileModelIntegrationTestSuite) SetupTest() {
	_, err := suite.database.CreateUser(testUser)
	suite.Assert().NoError(err)
}

func (suite *ProfileModelIntegrationTestSuite) TearDownTest() {
	globalTearDown(suite.database, suite.T())
}

func (suite *ProfileModelIntegrationTestSuite) TestCreateProfileIntegration() {
	// Given
	profile := testProfile

	// When
	createdProfile, err := suite.database.CreateProfile(profile)

	// Then
	suite.Require().NoError(err)
	suite.Assert().Equal(profile, createdProfile)
}

func (suite *ProfileModelIntegrationTestSuite) TestFindProfileByUserIDIntegration() {
	// Given
	profile := testProfile

	// When
	createdProfile, err := suite.database.CreateProfile(profile)
	suite.Require().NoError(err)
	foundProfile, err := suite.database.FindProfileByUserID(profile.UserID)

	// Then
	suite.Require().NoError(err)
	suite.Assert().Equal(profile, foundProfile)
	suite.Assert().Equal(profile, createdProfile)
}

func (suite *ProfileModelIntegrationTestSuite) TestFindProfileByUnknownUserIDIntegration() {
	// When
	profile, err := suite.database.FindProfileByUserID(domain.InvalidUserID())

	// Then
	suite.Require().Error(err)
	suite.Assert().Equal(domain.Profile{}, profile)
}

func (suite *ProfileModelIntegrationTestSuite) TestCreateProfileWithUnknownUserIntegration() {
	// Given
	profile := testProfile
	profile.UserID = domain.InvalidUserID()

	// When
	createdProfile, err := suite.database.CreateProfile(profile)

	// Then
	suite.Require().Error(err)
	suite.Assert().Equal(domain.Profile{}, createdProfile)
}

func (suite *ProfileModelIntegrationTestSuite) TestUpdateProfileIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *ProfileModelIntegrationTestSuite) TestTryUpdateProfileID() {
	suite.T().Skip("Not implemented")
}

func (suite *ProfileModelIntegrationTestSuite) TestTryUpdateProfileUserID() {
	suite.T().Skip("Not implemented")
}

package postgres

import (
	"back/src/core/domain"
	"github.com/stretchr/testify/suite"
)

type UserModelIntegrationTestSuite struct {
	suite.Suite
	database *Database
}

var testUser = domain.User{
	ID:       domain.UserID("f1ec7fce-9d1c-4cd4-a9e4-f2f0538f466f"),
	Email:    "test@mail.com",
	Password: "-CurrentUserPassword123-",
}

func (suite *DatabaseDomainBridgeTestSuite) TestUserToDomain() {
	//Given
	modelUser := User{
		ID:       testUser.ID.String(),
		Email:    testUser.Email,
		Password: testUser.Password,
	}

	//When
	domainUser := userToDomain(modelUser)

	//Then
	suite.Assert().Equal(modelUser.ID, domainUser.ID.String())
	suite.Assert().Equal(modelUser.Email, domainUser.Email)
	suite.Assert().Equal(modelUser.Password, domainUser.Password)
}

func (suite *DatabaseDomainBridgeTestSuite) TestUserToSession() {
	//Given
	domainUser := testUser

	//When
	modelUser := domainToUser(domainUser)

	//Then
	suite.Assert().Equal(testUser.ID.String(), modelUser.ID)
	suite.Assert().Equal(testUser.Email, modelUser.Email)
	suite.Assert().Equal(testUser.Password, modelUser.Password)
}

func (suite *DatabaseDomainBridgeTestSuite) TestUserDomainToModelToDomain() {
	//Given
	domainUser := testUser

	//When
	user := domainToUser(domainUser)
	userDomain := userToDomain(user)

	//Then
	suite.Assert().Equal(domainUser, userDomain)
}

func (suite *UserModelIntegrationTestSuite) SetupTest() {
}

func (suite *UserModelIntegrationTestSuite) TearDownTest() {
	globalTearDown(suite.database, suite.T())
}

func (suite *UserModelIntegrationTestSuite) TestCreateUserIntegration() {
	// Given
	user := testUser

	// When
	createdUser, err := suite.database.CreateUser(user)
	suite.Assert().NoError(err)
	foundUser, err := suite.database.FindUserByID(user.ID)

	// Then
	suite.Assert().NoError(err)
	suite.Assert().Equal(testUser, createdUser)
	suite.Assert().Equal(testUser, foundUser)
}

func (suite *UserModelIntegrationTestSuite) TestCreateUserWithoutIdIntegration() {
	// Given
	user := testUser
	user.ID = ""

	// When
	createdUser, err := suite.database.CreateUser(user)
	suite.Assert().Error(err)
	foundUser, err := suite.database.FindUserByID(user.ID)

	// Then
	suite.Assert().Error(err)
	suite.Assert().Equal(domain.User{}, createdUser)
	suite.Assert().Equal(domain.User{}, foundUser)
}

func (suite *UserModelIntegrationTestSuite) TestCreateUserWithDuplicateMailIntegration() {
	// Given
	user := testUser

	// When
	_, err := suite.database.CreateUser(user)
	user.ID = "b7c8e647-c859-4ebb-8b10-449fab8909d9"
	suite.Assert().NoError(err)
	_, err = suite.database.CreateUser(user)

	// Then
	suite.Assert().Error(err)
}

func (suite *UserModelIntegrationTestSuite) TestCreateUserWithDuplicateIdsIntegration() {
	// Given
	user := testUser

	// When
	_, err := suite.database.CreateUser(user)
	user.Email = "test2@mail.com"
	suite.Assert().NoError(err)
	_, err = suite.database.CreateUser(user)

	// Then
	suite.Assert().Error(err)
}

func (suite *UserModelIntegrationTestSuite) TestFindUserByIdIntegration() {
	// Given
	user := testUser

	// When
	createdUser, err := suite.database.CreateUser(user)
	suite.Assert().NoError(err)
	foundUser, err := suite.database.FindUserByID(user.ID)

	// Then
	suite.Assert().NoError(err)
	suite.Assert().Equal(testUser, createdUser)
	suite.Assert().Equal(testUser, foundUser)
}

func (suite *UserModelIntegrationTestSuite) TestFindUserByEmailIntegration() {
	// Given
	user := testUser

	// When
	createdUser, err := suite.database.CreateUser(user)
	suite.Assert().NoError(err)
	foundUser, err := suite.database.FindUserByEmail(user.Email)

	// Then
	suite.Assert().NoError(err)
	suite.Assert().Equal(testUser, createdUser)
	suite.Assert().Equal(testUser, foundUser)
}

func (suite *UserModelIntegrationTestSuite) TestCountUsersEmptyIntegration() {
	// When
	count := suite.database.CountUsers()

	// Then
	suite.Assert().Equal(int64(0), count)
}

func (suite *UserModelIntegrationTestSuite) TestCountUsersIntegration() {
	// Given
	user := testUser

	// When
	_, err := suite.database.CreateUser(user)
	suite.Assert().NoError(err)
	count := suite.database.CountUsers()

	// Then
	suite.Assert().Equal(int64(1), count)
}

func (suite *UserModelIntegrationTestSuite) TestUpdateUserMailIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *UserModelIntegrationTestSuite) TestUpdateUserPasswordIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *UserModelIntegrationTestSuite) TestUpdateUserMailAndPasswordIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *UserModelIntegrationTestSuite) TestUpdateNotExistingUserShouldThrowErrorIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *UserModelIntegrationTestSuite) TestDeleteUserIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *UserModelIntegrationTestSuite) TestDeleteNotExistingUserShouldThrowErrorIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *UserModelIntegrationTestSuite) TestRecreatingDeletedUser() {
	suite.T().Skip("Not implemented")
}

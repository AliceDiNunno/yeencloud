package postgres

import (
	"back/src/core/domain"
)

var testUser = domain.User{
	ID:       "User ID",
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

func (suite *DatabaseIntegrationTestSuite) TestCreateUserIntegration() {
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

func (suite *DatabaseIntegrationTestSuite) TestCreateUserWithoutIdIntegration() {
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

func (suite *DatabaseIntegrationTestSuite) TestCreateUserWithDuplicateMailIntegration() {
	// Given
	user := testUser

	// When
	_, err := suite.database.CreateUser(user)
	user.ID = "User ID 2"
	suite.Assert().NoError(err)
	_, err = suite.database.CreateUser(user)

	// Then
	suite.Assert().Error(err)
}

func (suite *DatabaseIntegrationTestSuite) TestCreateUserWithDuplicateIdsIntegration() {
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

func (suite *DatabaseIntegrationTestSuite) TestFindUserByIdIntegration() {
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

func (suite *DatabaseIntegrationTestSuite) TestFindUserByEmailIntegration() {
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

func (suite *DatabaseIntegrationTestSuite) TestCountUsersEmptyIntegration() {
	// When
	count := suite.database.CountUsers()

	// Then
	suite.Assert().Equal(int64(0), count)
}

func (suite *DatabaseIntegrationTestSuite) TestCountUsersIntegration() {
	// Given
	user := testUser

	// When
	_, err := suite.database.CreateUser(user)
	suite.Assert().NoError(err)
	count := suite.database.CountUsers()

	// Then
	suite.Assert().Equal(int64(1), count)
}

func (suite *DatabaseIntegrationTestSuite) TestUpdateUserMailIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *DatabaseIntegrationTestSuite) TestUpdateUserPasswordIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *DatabaseIntegrationTestSuite) TestUpdateUserMailAndPasswordIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *DatabaseIntegrationTestSuite) TestUpdateNotExistingUserShouldThrowErrorIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *DatabaseIntegrationTestSuite) TestDeleteUserIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *DatabaseIntegrationTestSuite) TestDeleteNotExistingUserShouldThrowErrorIntegration() {
	suite.T().Skip("Not implemented")
}

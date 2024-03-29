package postgres

import "back/src/core/domain"

func (suite *DatabaseDomainBridgeTestSuite) TestProfileToDomain() {
	//Given
	modelProfile := Profile{
		ID:       "ProfileID",
		UserID:   "UserID",
		Name:     "User Name",
		Language: "en-US",
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
	domainProfile := domain.Profile{
		ID:       "ProfileID",
		UserID:   "UserID",
		Name:     "User Name",
		Language: "en-US",
	}

	//When
	modelProfile := domainToProfile(domainProfile)

	//Then
	suite.Assert().Equal(domainProfile.ID.String(), modelProfile.ID)
	suite.Assert().Equal(domainProfile.UserID.String(), modelProfile.UserID)
	suite.Assert().Equal(domainProfile.Name, modelProfile.Name)
	suite.Assert().Equal(domainProfile.Language, modelProfile.Language)
}

func (suite *DatabaseDomainBridgeTestSuite) TestProfileDomainToModelToDomain() {
	//Given
	domainUser := domain.User{
		ID:       "User ID",
		Email:    "test@mail.com",
		Password: "-CurrentUserPassword123-",
	}

	//When
	user := domainToUser(domainUser)
	userDomain := userToDomain(user)

	//Then
	suite.Assert().Equal(domainUser, userDomain)
}

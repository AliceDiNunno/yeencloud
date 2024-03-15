package postgres

import "back/src/core/domain"

func (suite *DatabaseDomainBridgeTestSuite) TestOrganizationToDomain() {
	//Given
	modelOrganization := Organization{
		ID:          "OrganizationID",
		Slug:        "Slug",
		Name:        "Name",
		Description: "Description",
	}

	//When
	domainOrganization := organizationToDomain(modelOrganization)

	//Then
	suite.Assert().Equal(modelOrganization.ID, domainOrganization.ID.String())
	suite.Assert().Equal(modelOrganization.Slug, domainOrganization.Slug)
	suite.Assert().Equal(modelOrganization.Name, domainOrganization.Name)
	suite.Assert().Equal(modelOrganization.Description, domainOrganization.Description)
}

func (suite *DatabaseDomainBridgeTestSuite) TestDomainToOrganization() {
	//Given
	domainOrganization := domain.Organization{
		ID:          "OrganizationID",
		Slug:        "Slug",
		Name:        "Name",
		Description: "Description",
	}

	//When
	modelOrganization := domainToOrganization(domainOrganization)

	//Then
	suite.Assert().Equal(domainOrganization.ID.String(), modelOrganization.ID)
	suite.Assert().Equal(domainOrganization.Slug, modelOrganization.Slug)
	suite.Assert().Equal(domainOrganization.Name, modelOrganization.Name)
	suite.Assert().Equal(domainOrganization.Description, modelOrganization.Description)
}

func (suite *DatabaseDomainBridgeTestSuite) TestOrganizationDomainToModelToDomain() {
	//Given
	domainUser := domain.Organization{
		ID:          "OrganizationID",
		Slug:        "Slug",
		Name:        "Name",
		Description: "Description",
	}

	//When
	user := domainToOrganization(domainUser)
	userDomain := organizationToDomain(user)

	//Then
	suite.Assert().Equal(domainUser, userDomain)
}

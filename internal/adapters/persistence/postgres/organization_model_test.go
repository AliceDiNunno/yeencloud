package postgres

import (
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/stretchr/testify/suite"
)

type OrganizationModelIntegrationTestSuite struct {
	suite.Suite
	database *Database
}

var testOrganization = domain.Organization{
	ID:          domain.OrganizationID("d0348c8c-1cff-414d-b8f3-3f3d6a826918"),
	Slug:        "test-organization",
	Name:        "Test Organization",
	Description: "This is an organization for testing purposes.",
}

func (suite *DatabaseDomainBridgeTestSuite) TestOrganizationToDomain() {
	// Given
	modelOrganization := Organization{
		ID:          testOrganization.ID.String(),
		Slug:        testOrganization.Slug,
		Name:        testOrganization.Name,
		Description: testOrganization.Description,
	}

	// When
	domainOrganization := organizationToDomain(modelOrganization)

	// Then
	suite.Assert().Equal(modelOrganization.ID, domainOrganization.ID.String())
	suite.Assert().Equal(modelOrganization.Slug, domainOrganization.Slug)
	suite.Assert().Equal(modelOrganization.Name, domainOrganization.Name)
	suite.Assert().Equal(modelOrganization.Description, domainOrganization.Description)
}

func (suite *DatabaseDomainBridgeTestSuite) TestDomainToOrganization() {
	// Given
	domainOrganization := testOrganization

	// When
	modelOrganization := domainToOrganization(domainOrganization)

	// Then
	suite.Assert().Equal(testOrganization.ID.String(), modelOrganization.ID)
	suite.Assert().Equal(testOrganization.Slug, modelOrganization.Slug)
	suite.Assert().Equal(testOrganization.Name, modelOrganization.Name)
	suite.Assert().Equal(testOrganization.Description, modelOrganization.Description)
}

func (suite *DatabaseDomainBridgeTestSuite) TestOrganizationDomainToModelToDomain() {
	// Given
	domainOrganization := testOrganization

	// When
	organization := domainToOrganization(domainOrganization)
	organizationDomain := organizationToDomain(organization)

	// Then
	suite.Assert().Equal(domainOrganization, organizationDomain)
}

func (suite *OrganizationModelIntegrationTestSuite) SetupTest() {

}

func (suite *OrganizationModelIntegrationTestSuite) TearDownTest() {
	globalTearDown(suite.database, suite.T())
}

func (suite *OrganizationModelIntegrationTestSuite) TestCreateOrganizationIntegration() {
	// Given
	organization := testOrganization

	// When
	createdOrganization, err := suite.database.CreateOrganization(organization)
	suite.Assert().NoError(err)
	foundOrganization, err := suite.database.FindOrganizationByID(organization.ID)

	// Then
	suite.Assert().NoError(err)
	suite.Assert().Equal(testOrganization, createdOrganization)
	suite.Assert().Equal(testOrganization, foundOrganization)
}

func (suite *OrganizationModelIntegrationTestSuite) TestCreateOrganizationWithoutIdIntegration() {
	// Given
	organization := testOrganization
	organization.ID = ""

	// When
	createdOrganization, err := suite.database.CreateOrganization(organization)
	suite.Assert().Error(err)
	foundOrganization, err := suite.database.FindOrganizationByID(organization.ID)

	// Then
	suite.Assert().Error(err)
	suite.Assert().Equal(domain.Organization{}, createdOrganization)
	suite.Assert().Equal(domain.Organization{}, foundOrganization)
}

func (suite *OrganizationModelIntegrationTestSuite) TestCreateOrganizationWithDuplicateIdsIntegration() {
	// Given
	organization := testOrganization

	// When
	_, err := suite.database.CreateOrganization(organization)
	suite.Assert().NoError(err)
	_, err = suite.database.CreateOrganization(organization)

	// Then
	suite.Assert().Error(err)
}

func (suite *OrganizationModelIntegrationTestSuite) TestCreateOrganizationWithDuplicateSlugsIntegration() {
	// Given
	organization := testOrganization

	// When
	_, err := suite.database.CreateOrganization(organization)
	organization.ID = "930f9b50-cfe1-4514-8c92-6ad7e89665ab"
	suite.Assert().NoError(err)
	_, err = suite.database.CreateOrganization(organization)

	// Then
	suite.Assert().Error(err)
}

func (suite *OrganizationModelIntegrationTestSuite) TestFindOrganizationByIdIntegration() {
	// Given
	organization := testOrganization

	// When
	createdOrganization, err := suite.database.CreateOrganization(organization)
	suite.Assert().NoError(err)
	foundOrganization, err := suite.database.FindOrganizationByID(organization.ID)

	// Then
	suite.Assert().NoError(err)
	suite.Assert().Equal(testOrganization, createdOrganization)
	suite.Assert().Equal(testOrganization, foundOrganization)
}

func (suite *OrganizationModelIntegrationTestSuite) TestFindOrganizationBySlugIntegration() {
	// Given
	organization := testOrganization

	// When
	createdOrganization, err := suite.database.CreateOrganization(organization)
	suite.Assert().NoError(err)
	foundOrganization, err := suite.database.FindOrganizationBySlug(organization.Slug)

	// Then
	suite.Assert().NoError(err)
	suite.Assert().Equal(testOrganization, createdOrganization)
	suite.Assert().Equal(testOrganization, foundOrganization)
}

func (suite *OrganizationModelIntegrationTestSuite) TestUpdateOrganizationFieldsIntegration() {
	// Given
	organization := testOrganization
	update := domain.UpdateOrganization{
		Name:        "Updated Test Organization",
		Description: "Updated description for the organization.",
	}

	// When
	createdOrganization, err := suite.database.CreateOrganization(organization)
	suite.Assert().NoError(err)
	updatedOrganization, err := suite.database.UpdateOrganization(organization.ID, update)

	// Then
	suite.Assert().NoError(err)
	suite.Assert().Equal(testOrganization, createdOrganization)
	suite.Assert().Equal(updatedOrganization.Name, update.Name)
	suite.Assert().Equal(updatedOrganization.Description, update.Description)
	suite.Assert().Equal(updatedOrganization.ID, createdOrganization.ID)
	suite.Assert().Equal(updatedOrganization.Slug, createdOrganization.Slug)
}

func (suite *OrganizationModelIntegrationTestSuite) TestUpdateNotExistingOrganizationShouldThrowErrorIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *OrganizationModelIntegrationTestSuite) TestDeleteOrganizationIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *OrganizationModelIntegrationTestSuite) TestDeleteNotExistingOrganizationShouldThrowErrorIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *OrganizationModelIntegrationTestSuite) TestRecreatingDeletedOrganizationIntegration() {
	// keeping slug
	suite.T().Skip("Not implemented")
}

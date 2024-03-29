package postgres

import (
	"back/src/core/domain"
	"github.com/stretchr/testify/suite"
	"time"
)

type SessionModelIntegrationTestSuite struct {
	suite.Suite
	database *Database
}

var testIP = "127.0.0.1"
var testValidSession = domain.Session{
	Token:    "8784e1bb-bf81-4ca1-8ad3-97eddb25972a",
	ExpireAt: time.Now().AddDate(0, 0, 1).UnixMilli(),
	IP:       testIP,
	UserID:   testUser.ID,
}

var testExpiredSession = domain.Session{
	Token:    "20b2b56e-b48d-4fdb-9ac9-099ec9a3925a",
	ExpireAt: 0,
	IP:       testIP,
	UserID:   testUser.ID,
}

var testNoUserSession = domain.Session{
	Token:    "7d9c499f-ac5f-4968-9ba3-548abcc6995a",
	ExpireAt: time.Now().AddDate(0, 0, 1).UnixMilli(),
	IP:       testIP,
	UserID:   domain.InvalidUserID(),
}

func (suite *DatabaseDomainBridgeTestSuite) TestSessionToDomain() {
	//Given
	modelSession := Session{
		Token:    testValidSession.Token,
		IP:       testValidSession.IP,
		ExpireAt: time.UnixMilli(testValidSession.ExpireAt),
		UserID:   testValidSession.UserID.String(),
	}

	//When
	domainSession := sessionToDomain(modelSession)

	//Then
	suite.Assert().Equal(modelSession.Token, domainSession.Token)
	suite.Assert().Equal(modelSession.IP, domainSession.IP)
	suite.Assert().Equal(modelSession.ExpireAt.Unix(), domainSession.ExpireAt)
	suite.Assert().Equal(modelSession.UserID, domainSession.UserID.String())
}

func (suite *DatabaseDomainBridgeTestSuite) TestDomainToSession() {
	//Given
	domainSession := testValidSession

	//When
	modelSession := domainToSession(domainSession)

	//Then
	suite.Assert().Equal(testValidSession.Token, modelSession.Token)
	suite.Assert().Equal(testValidSession.IP, modelSession.IP)
	suite.Assert().Equal(testValidSession.ExpireAt, modelSession.ExpireAt.Unix())
	suite.Assert().Equal(testValidSession.UserID.String(), modelSession.UserID)
}

func (suite *DatabaseDomainBridgeTestSuite) TestSessionDomainToModelToDomain() {
	//Given
	domainSession := testValidSession

	//When
	session := domainToSession(domainSession)
	sessionDomain := sessionToDomain(session)

	//Then
	suite.Assert().Equal(testValidSession, sessionDomain)
}

func (suite *SessionModelIntegrationTestSuite) SetupTest() {
	_, err := suite.database.CreateUser(testUser)
	suite.Assert().NoError(err)
}

func (suite *SessionModelIntegrationTestSuite) TearDownTest() {
	globalTearDown(suite.database, suite.T())
}

func (suite *SessionModelIntegrationTestSuite) TestCreateSessionIntegration() {
	// Given
	session := testValidSession

	// When
	createdSession, err := suite.database.CreateSession(session)
	suite.Assert().NoError(err)
	foundSession, err := suite.database.FindSessionByToken(session.Token)

	// Then
	suite.Assert().NoError(err)
	suite.Assert().Equal(testValidSession, createdSession)
	suite.Assert().Equal(testValidSession, foundSession)
}

func (suite *SessionModelIntegrationTestSuite) TestCreateMultipleSessionsIntegration() {
	// Given
	session1 := testValidSession
	session2 := testValidSession
	session2.Token = "3a3eae06-3229-4578-a5d7-ec93b7f37b27"

	// When
	createdSession1, err := suite.database.CreateSession(session1)
	suite.Assert().NoError(err)
	foundSession1, err := suite.database.FindSessionByToken(session1.Token)
	suite.Assert().NoError(err)

	createdSession2, err := suite.database.CreateSession(session2)
	suite.Assert().NoError(err)
	foundSession2, err := suite.database.FindSessionByToken(session2.Token)

	// Then
	suite.Assert().NoError(err)
	suite.Assert().Equal(testValidSession, createdSession1)
	suite.Assert().Equal(testValidSession, foundSession1)
	suite.Assert().Equal(testValidSession, createdSession2)
	suite.Assert().Equal(testValidSession, foundSession2)
}

func (suite *SessionModelIntegrationTestSuite) TestCreateSessionWithInvalidTokenIntegration() {
	// Given
	session := testValidSession
	session.Token = ""

	// When
	createdSession, err := suite.database.CreateSession(session)
	suite.Assert().Error(err)
	foundSession, err := suite.database.FindSessionByToken(session.Token)

	// Then
	suite.Assert().Error(err)
	suite.Assert().Equal(domain.Session{}, createdSession)
	suite.Assert().Equal(domain.Session{}, foundSession)
}

func (suite *SessionModelIntegrationTestSuite) TestCreateExpiredSessionIntegration() {
	// Given
	session := testExpiredSession

	// When
	createdSession, err := suite.database.CreateSession(session)
	suite.Assert().NoError(err)
	foundSession, err := suite.database.FindSessionByToken(session.Token)

	// Then
	suite.Assert().Error(err)
	suite.Assert().Equal(testExpiredSession, createdSession)
	suite.Assert().Equal(domain.Session{}, foundSession)
}

func (suite *SessionModelIntegrationTestSuite) TestFindSessionByTokenIntegration() {
	// Given
	session := testValidSession
	createdSession, err := suite.database.CreateSession(session)
	suite.Assert().NoError(err)

	// When
	foundSession, err := suite.database.FindSessionByToken(session.Token)

	// Then
	suite.Assert().NoError(err)
	suite.Assert().Equal(testValidSession, createdSession)
	suite.Assert().Equal(testValidSession, foundSession)
}

func (suite *SessionModelIntegrationTestSuite) TestFindSessionByTokenNotFoundIntegration() {
	// Given
	session := testValidSession
	_, err := suite.database.CreateSession(session)
	suite.Assert().NoError(err)

	// When
	foundSession, err := suite.database.FindSessionByToken("7cea29f2-053b-410c-9825-81a915ea3919")

	// Then
	suite.Assert().Error(err)
	suite.Assert().Equal(domain.Session{}, foundSession)
}

func (suite *SessionModelIntegrationTestSuite) TestCreateSessionWithInvalidUserIDIntegration() {
	// Given
	session := testNoUserSession

	// When
	createdSession, err := suite.database.CreateSession(session)
	suite.Assert().Error(err)
	foundSession, err := suite.database.FindSessionByToken(session.Token)

	// Then
	suite.Assert().Error(err)
	suite.Assert().Equal(domain.Session{}, createdSession)
	suite.Assert().Equal(domain.Session{}, foundSession)
}

func (suite *SessionModelIntegrationTestSuite) TestTryToUpdateSessionIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *SessionModelIntegrationTestSuite) TestDeleteSessionByTokenIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *SessionModelIntegrationTestSuite) TestDeleteNonExistingSessionByTokenIntegration() {
	suite.T().Skip("Not implemented")
}

func (suite *SessionModelIntegrationTestSuite) TestDeleteAllUsersSessionIntegration() {
	suite.T().Skip("Not implemented")
}

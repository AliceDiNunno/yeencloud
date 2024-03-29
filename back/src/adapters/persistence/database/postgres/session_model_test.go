package postgres

import (
	"back/src/core/domain"
	"time"
)

func (suite *DatabaseDomainBridgeTestSuite) TestSessionToDomain() {
	//Given
	modelSession := Session{
		Token:    "Session Token",
		IP:       "127.0.0.1",
		ExpireAt: time.Now(),
		UserID:   "UserID",
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
	domainSession := domain.Session{
		Token:    "Session Token",
		IP:       "127.0.0.1",
		ExpireAt: time.Now().Unix(),
		UserID:   "UserID",
	}

	//When
	modelSession := domainToSession(domainSession)

	//Then
	suite.Assert().Equal(domainSession.Token, modelSession.Token)
	suite.Assert().Equal(domainSession.IP, modelSession.IP)
	suite.Assert().Equal(domainSession.ExpireAt, modelSession.ExpireAt.Unix())
	suite.Assert().Equal(domainSession.UserID.String(), modelSession.UserID)
}

func (suite *DatabaseDomainBridgeTestSuite) TestSessionDomainToModelToDomain() {
	//Given
	domainSession := domain.Session{
		Token:    "Session Token",
		IP:       "127.0.0.1",
		ExpireAt: time.Now().Unix(),
		UserID:   "UserID",
	}

	//When
	session := domainToSession(domainSession)
	sessionDomain := sessionToDomain(session)

	//Then
	suite.Assert().Equal(domainSession, sessionDomain)
}

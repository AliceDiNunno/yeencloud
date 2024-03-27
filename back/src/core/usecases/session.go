package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// #YC-21 TODO: should a session be in the usecases or the http layer?
func (self UCs) CreateSession(auditID domain.AuditID, newSessionRequest domain.NewSession) (domain.Session, *domain.ErrorDescription) {
	self.i.Auditer.AddStep(auditID, newSessionRequest.Secure())

	// #YC-3 TODO: implement OTP
	us, err := self.i.Persistence.User.FindUserByEmail(newSessionRequest.Email)

	if err != nil {
		return domain.Session{}, &domain.ErrorUserNotFound
	}

	if bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(newSessionRequest.Password)) != nil {
		return domain.Session{}, &domain.ErrorUserNotFound
	}

	sessionToken := uuid.New().String()
	// #YC-18 TODO: expiration time should be configurable
	expiration := time.Now().Add(365 * 24 * time.Hour)
	newSession := domain.Session{
		Token:    sessionToken,
		ExpireAt: expiration.Unix(),
		IP:       newSessionRequest.Origin,
		UserID:   us.ID,
	}

	session, err := self.i.Persistence.Session.CreateSession(newSession)
	if err != nil {
		return domain.Session{}, nil
	}

	return session, nil
}

func (self UCs) GetSessionByToken(auditID domain.AuditID, token string) (domain.Session, *domain.ErrorDescription) {
	self.i.Auditer.AddStep(auditID)

	// #YC-20 TODO: this should check if the user still exists and if the session is still valid
	session, err := self.i.Persistence.Session.FindSessionByToken(token)
	if err != nil {
		return domain.Session{}, &domain.ErrorSessionNotFound
	}

	return session, nil
}

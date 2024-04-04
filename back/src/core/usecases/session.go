package usecases

import (
	"time"

	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SessionUsecases interface {
	CreateSession(auditID domain.AuditID, user domain.NewSession) (domain.Session, *domain.ErrorDescription)

	GetSessionByToken(auditID domain.AuditID, token string) (domain.Session, *domain.ErrorDescription)
}

func (self UCs) CreateSession(auditID domain.AuditID, newSessionRequest domain.NewSession) (domain.Session, *domain.ErrorDescription) {
	step := self.i.Trace.AddStep(auditID, newSessionRequest.Secure())

	// #YC-3 TODO: implement OTP
	us, err := self.i.Persistence.FindUserByEmail(newSessionRequest.Email)

	if err != nil {
		self.i.Trace.EndStep(auditID, step)
		return domain.Session{}, &domain.ErrorUserNotFound
	}

	if bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(newSessionRequest.Password)) != nil {
		self.i.Trace.Log(auditID, step).WithLevel(domain.LogLevelWarn).WithField(domain.LogFieldSessionRequestMail, newSessionRequest.Email).Msg("User tried to login with wrong password")
		self.i.Trace.EndStep(auditID, step)
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

	session, err := self.i.Persistence.CreateSession(newSession)
	if err != nil {
		self.i.Trace.EndStep(auditID, step)
		return domain.Session{}, nil
	}

	self.i.Trace.EndStep(auditID, step)
	return session, nil
}

func (self UCs) GetSessionByToken(auditID domain.AuditID, token string) (domain.Session, *domain.ErrorDescription) {
	step := self.i.Trace.AddStep(auditID)

	// #YC-20 TODO: this should check if the user still exists and if the session is still valid
	session, err := self.i.Persistence.FindSessionByToken(token)
	if err != nil {
		self.i.Trace.EndStep(auditID, step)
		return domain.Session{}, &domain.ErrorSessionNotFound
	}

	self.i.Trace.EndStep(auditID, step)
	return session, nil
}

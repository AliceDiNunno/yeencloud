package usecases

import (
	"time"

	"github.com/AliceDiNunno/yeencloud/internal/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SessionUsecases interface {
	CreateSession(auditID domain.AuditTraceID, user domain.NewSession) (domain.Session, *domain.ErrorDescription)

	GetSessionByToken(auditID domain.AuditTraceID, token string) (domain.Session, *domain.ErrorDescription)
}

func (self UCs) createSession(auditID domain.AuditTraceID, origin string, us domain.User) (domain.Session, *domain.ErrorDescription) {
	step := self.i.Trace.AddStep(auditID, audit.DefaultSkip)

	sessionToken := uuid.New().String()
	// #YC-18 TODO: expiration time should be configurable
	expiration := time.Now().Add(365 * 24 * time.Hour)
	newSession := domain.Session{
		Token:    sessionToken,
		ExpireAt: expiration.Unix(),
		IP:       origin,
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

func (self UCs) CreateSession(auditID domain.AuditTraceID, newSessionRequest domain.NewSession) (domain.Session, *domain.ErrorDescription) {
	step := self.i.Trace.AddStep(auditID, audit.DefaultSkip, newSessionRequest.Secure())

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

	session, derr := self.createSession(auditID, newSessionRequest.Origin, us)
	if derr != nil {
		self.i.Trace.EndStep(auditID, step)
		return domain.Session{}, derr
	}

	self.i.Trace.EndStep(auditID, step)
	return session, nil
}

func (self UCs) GetSessionByToken(auditID domain.AuditTraceID, token string) (domain.Session, *domain.ErrorDescription) {
	step := self.i.Trace.AddStep(auditID, audit.DefaultSkip)

	// #YC-20 TODO: this should check if the user still exists and if the session is still valid
	session, err := self.i.Persistence.FindSessionByToken(token)
	if err != nil {
		self.i.Trace.EndStep(auditID, step)
		return domain.Session{}, &domain.ErrorSessionNotFound
	}

	self.i.Trace.EndStep(auditID, step)
	return session, nil
}

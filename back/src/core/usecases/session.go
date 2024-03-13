package usecases

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// #YC-21 TODO: should a session be in the usecases or the http layer?
func (i interactor) CreateSession(auditID domain.AuditID, newSessionRequest requests.NewSession) (domain.Session, *domain.ErrorDescription) {
	i.auditer.AddStep(auditID, newSessionRequest.Secure())

	// #YC-3 TODO: implement OTP
	us, err := i.userRepo.FindUserByEmail(newSessionRequest.Email)

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
		IP:       newSessionRequest.IP,
		UserID:   us.ID,
	}

	session, err := i.sessionRepo.CreateSession(newSession)
	if err != nil {
		return domain.Session{}, nil
	}

	return session, nil
}

func (i interactor) GetSessionByToken(auditID domain.AuditID, token string) (domain.Session, *domain.ErrorDescription) {
	i.auditer.AddStep(auditID)

	// #YC-20 TODO: this should check if the user still exists and if the session is still valid
	session, err := i.sessionRepo.FindSessionByToken(token)
	if err != nil {
		return domain.Session{}, &domain.ErrorSessionNotFound
	}

	return session, nil
}

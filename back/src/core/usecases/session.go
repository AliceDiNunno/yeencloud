package usecases

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// TODO: should a session be in the usecases or the http layer?
func (i interactor) CreateSession(newSessionRequest requests.NewSession) (domain.Session, *domain.ErrorDescription) {
	//TODO: implement OTP
	us, err := i.userRepo.FindUserByEmail(newSessionRequest.Email)

	if err != nil {
		return domain.Session{}, &domain.ErrorUserNotFound
	}

	if bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(newSessionRequest.Password)) != nil {
		return domain.Session{}, &domain.ErrorUserNotFound
	}

	sessionToken := uuid.New().String()
	//TODO: expiration time should be configurable
	expiration := time.Now().Add(365 * 24 * time.Hour)
	newSession := domain.Session{
		Token:    sessionToken,
		ExpireAt: expiration.Unix(),
		IP:       newSessionRequest.IP,
		UserID:   us.ID,
	}

	session, err := i.sessionRepository.CreateSession(newSession)
	if err != nil {
		return domain.Session{}, nil
	}

	return session, nil
}

func (i interactor) GetSessionByToken(token string) (domain.Session, *domain.ErrorDescription) {
	//TODO: this should check if the user still exists and if the session is still valid

	session, err := i.sessionRepository.FindSessionByToken(token)
	if err != nil {
		return domain.Session{}, &domain.ErrorSessionNotFound
	}

	return session, nil
}

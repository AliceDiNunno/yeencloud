package postgres

import (
	"back/src/core/domain"
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model

	Token    string `gorm:"primary_key"`
	IP       string
	ExpireAt time.Time

	UserID string
	User   User
}

func (db *Database) CreateSession(session domain.Session) (domain.Session, error) {
	sessionToCreate := domainToSession(session)

	result := db.engine.Create(&sessionToCreate)

	if result.Error != nil {
		return domain.Session{}, result.Error
	}

	return sessionToDomain(sessionToCreate), nil
}

func (db *Database) FindSessionByToken(token string) (domain.Session, error) {
	var session Session

	result := db.engine.Where("token = ?", token).First(&session)

	if result.Error != nil {
		return domain.Session{}, result.Error
	}

	return sessionToDomain(session), nil
}

func sessionToDomain(session Session) domain.Session {
	return domain.Session{
		Token:    session.Token,
		ExpireAt: session.ExpireAt.Unix(),
		IP:       session.IP,
		UserID:   domain.UserID(session.UserID),
	}
}

func domainToSession(session domain.Session) Session {
	return Session{
		Token:    session.Token,
		IP:       session.IP,
		ExpireAt: time.Unix(session.ExpireAt, 0),
		UserID:   session.UserID.String(),
	}
}
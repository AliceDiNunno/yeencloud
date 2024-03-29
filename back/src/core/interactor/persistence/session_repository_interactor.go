package persistence

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type SessionRepository interface {
	CreateSession(session domain.Session) (domain.Session, error)

	FindSessionByToken(token string) (domain.Session, error)
}

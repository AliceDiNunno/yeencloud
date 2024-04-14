package persistence

import "github.com/AliceDiNunno/yeencloud/internal/core/domain"

type SessionRepository interface {
	CreateSession(session domain.Session) (domain.Session, error)

	FindSessionByToken(token string) (domain.Session, error)
}

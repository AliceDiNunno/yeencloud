package persistence

import "github.com/AliceDiNunno/yeencloud/internal/core/domain"

type TokenRepository interface {
	CreateToken(session domain.Token) (domain.Token, error)
	FindToken(mail string, token string, tokenType domain.TokenType) (domain.Token, error)
	InvalidateToken(token domain.TokenID) error
}

package postgres

import (
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model

	ID string `gorm:"primary_key;unique;not null;default:null;<-:create"`

	CreatedAt int64
	ExpireAt  int64

	UserID string
	User   User

	Token string

	Type string
}

func (db *Database) CreateToken(domainToken domain.Token) (domain.Token, error) {
	token := domainToToken(domainToken)

	result := db.engine.Create(&token)

	if result.Error != nil {
		return domain.Token{}, sqlerr(result.Error)
	}

	return tokenToDomain(token), nil
}

func (db *Database) FindToken(mail string, token string, tokenType domain.TokenType) (domain.Token, error) {
	var tokenModel Token

	result := db.engine.Preload("User").Joins("LEFT JOIN users on users.id = user_id").Where("token = ? AND type = ? AND users.email = ?", token, tokenType, mail).First(&tokenModel)

	if result.Error != nil {
		return domain.Token{}, sqlerr(result.Error)
	}

	return tokenToDomain(tokenModel), nil
}

func (db *Database) InvalidateToken(token domain.TokenID) error {
	result := db.engine.Where("id = ?", token).Delete(&Token{})

	if result.Error != nil {
		return sqlerr(result.Error)
	}

	return nil
}

func domainToToken(token domain.Token) Token {
	return Token{
		ID:        token.ID.String(),
		UserID:    token.User.ID.String(),
		Token:     token.Token,
		CreatedAt: token.CreatedAt,
		ExpireAt:  token.ExpireAt,
		Type:      token.Type.String(),
	}
}

func tokenToDomain(token Token) domain.Token {
	return domain.Token{
		ID:        domain.TokenID(token.ID),
		User:      userToDomain(token.User),
		Token:     token.Token,
		CreatedAt: token.CreatedAt,
		ExpireAt:  token.ExpireAt,
		Type:      domain.TokenType(token.Type),
	}
}

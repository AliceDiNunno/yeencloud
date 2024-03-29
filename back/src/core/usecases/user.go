package usecases

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (i interactor) CreateUser(newUser requests.NewUser, profileLanguage string) (domain.User, *domain.ErrorDescription) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	userToCreate := domain.User{
		ID:       uuid.New().String(),
		Email:    newUser.Email,
		Password: string(hashedPassword),
	}

	user, err := i.userRepo.CreateUser(userToCreate)

	if err != nil {
		log.Err(err).Str("mail", newUser.Email).Msg("Error creating user")
	}

	profileToCreate := domain.Profile{
		UserID:   user.ID,
		Name:     newUser.Name,
		Language: profileLanguage,
	}

	_, err = i.profileRepo.CreateProfile(profileToCreate)

	if err != nil {
		log.Err(err).Str("mail", newUser.Email).Msg("Error creating profile for user")
	}

	log.Info().Str("mail", newUser.Email).Msg("User created")
	return domain.User{}, &domain.ErrorNoMethod
}

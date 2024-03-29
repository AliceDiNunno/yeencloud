package usecases

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (i interactor) CreateUser(newUser requests.NewUser, profileLanguage string) (domain.Profile, *domain.ErrorDescription) {
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
	return profileToCreate, &domain.ErrorNoMethod
}

func (i interactor) GetUserByID(id string) (domain.User, *domain.ErrorDescription) {
	user, err := i.userRepo.FindUserByID(uuid.MustParse(id))

	if err != nil {
		log.Err(err).Str("id", id).Msg("Error finding user")
		return domain.User{}, &domain.ErrorUserNotFound
	}

	return user, nil
}

func (i interactor) GetProfileByUserID(id string) (domain.Profile, *domain.ErrorDescription) {
	profile, err := i.profileRepo.FindProfileByUserID(uuid.MustParse(id))

	//TODO: this should never happen, a profile should be created if it ever is missing (while also reporting the error so it can be investigated)
	if err != nil {
		log.Err(err).Str("id", id).Msg("Error finding profile")
		return domain.Profile{}, &domain.ErrorProfileNotFound
	}

	return profile, nil
}

package usecases

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (i interactor) newUserID() domain.UserID {
	return domain.UserID(uuid.New().String())
}

func (i interactor) CreateUser(auditID domain.AuditID, newUser requests.NewUser, profileLanguage string) (domain.Profile, *domain.ErrorDescription) {
	i.auditer.AddStep(auditID, newUser.Secure())

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	userToCreate := domain.User{
		ID:       i.newUserID(),
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

	msg := i18n.NewLocalizer(i.translator, profileLanguage)

	localizedDescription, _, _ := msg.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID: domain.DefaultOrganizationDescription,
		TemplateData: map[string]interface{}{
			domain.DefaultOrganizationDescriptionKey: newUser.Name,
		},
	})

	organizationToCreate := requests.NewOrganization{
		Name:        newUser.Name,
		Description: localizedDescription,
	}

	_, derr := i.CreateOrganization(auditID, user.ID, organizationToCreate)

	if derr != nil {
		return domain.Profile{}, derr
	}

	log.Info().Str("mail", newUser.Email).Msg("User created")
	return profileToCreate, nil
}

func (i interactor) GetUserByID(auditID domain.AuditID, id domain.UserID) (domain.User, *domain.ErrorDescription) {
	user, err := i.userRepo.FindUserByID(id)

	if err != nil {
		log.Err(err).Str("id", id.String()).Msg("Error finding user")
		return domain.User{}, &domain.ErrorUserNotFound
	}

	return user, nil
}

func (i interactor) GetProfileByUserID(auditID domain.AuditID, userID domain.UserID) (domain.Profile, *domain.ErrorDescription) {
	i.auditer.AddStep(auditID)
	profile, err := i.profileRepo.FindProfileByUserID(userID)

	// #YC-22 TODO: this should never happen, a profile should be created if it ever is missing (while also reporting the error so it can be investigated)
	if err != nil {
		log.Err(err).Str("id", userID.String()).Msg("Error finding profile")
		return domain.Profile{}, &domain.ErrorProfileNotFound
	}

	return profile, nil
}

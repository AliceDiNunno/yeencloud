package usecases

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (i interactor) CreateOrganization(userID domain.UserID, newOrganization requests.NewOrganization) (domain.Organization, *domain.ErrorDescription) {
	organizationToCreate := domain.Organization{
		ID:          domain.OrganizationID(uuid.New().String()),
		Slug:        newOrganization.Name, //TODO: generate unique slug
		Name:        newOrganization.Name,
		Description: newOrganization.Description,
	}

	organization, err := i.organizationRepo.CreateOrganization(organizationToCreate)

	if err != nil {
		log.Err(err).Str("id", userID.String()).Msg("Error creating organization for user")
	}

	err = i.organizationUserRepo.LinkUserToOrganization(userID, organization.ID, "admin")

	if err != nil {
		log.Err(err).Str("id", userID.String()).Msg("Error linking user to organization")
	}

	return organization, nil
}

func (i interactor) GetOrganizationsByUserID(userID domain.UserID) ([]domain.OrganizationMember, *domain.ErrorDescription) {
	organizations, err := i.organizationUserRepo.GetUserOrganizationsByUserID(userID)

	if err != nil {
		log.Err(err).Str("id", userID.String()).Msg("Error getting organizations for user")
		return nil, &domain.ErrorUnableToGetUserOrganizations
	}

	return organizations, nil
}

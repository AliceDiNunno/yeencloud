package usecases

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (i interactor) CreateOrganization(auditID domain.AuditID, profileID domain.ProfileID, newOrganization requests.NewOrganization) (domain.Organization, *domain.ErrorDescription) {
	i.auditer.AddStep(auditID, newOrganization)

	organizationToCreate := domain.Organization{
		ID:          domain.OrganizationID(uuid.New().String()),
		Slug:        newOrganization.Name, // #YC-17 TODO: generate unique slug
		Name:        newOrganization.Name,
		Description: newOrganization.Description,
	}

	organization, err := i.organizationRepo.CreateOrganization(organizationToCreate)

	if err != nil {
		log.Err(err).Str("id", profileID.String()).Msg("Error creating organization for user")
	}

	err = i.organizationUserRepo.LinkProfileToOrganization(profileID, organization.ID, "admin")

	if err != nil {
		log.Err(err).Str("id", profileID.String()).Msg("Error linking user to organization")
	}

	return organization, nil
}

func (i interactor) GetOrganizationsByProfileID(auditID domain.AuditID, profileID domain.ProfileID) ([]domain.OrganizationMember, *domain.ErrorDescription) {
	i.auditer.AddStep(auditID, profileID)

	organizations, err := i.organizationUserRepo.GetProfileOrganizationsByProfileID(profileID)

	if err != nil {
		log.Err(err).Str("id", profileID.String()).Msg("Error getting organizations for user")
		return nil, &domain.ErrorUnableToGetUserOrganizations
	}

	return organizations, nil
}

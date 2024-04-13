package usecases

import (
	"crypto/rand"
	"io"
	"os"

	"github.com/AliceDiNunno/yeencloud/src/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
)

type TokenUsecases interface {
	ValidateMail(auditID domain.AuditTraceID, validation domain.ValidateMail) (domain.Session, *domain.ErrorDescription)
}

func (self UCs) generateToken() string {
	if os.Getenv("ENV") == "dev" {
		return "123456"
	}

	// generate a 6 digit token
	table := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	max := 6
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func (self UCs) ValidateMail(auditID domain.AuditTraceID, validation domain.ValidateMail) (domain.Session, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip)

	token, err := self.i.Persistence.FindToken(validation.Email, validation.Token, domain.TokenTypeVerifyEmail)

	if err != nil {
		return domain.Session{}, &domain.ErrorTokenNotFound
	}

	err = self.i.Persistence.InvalidateToken(token.ID)

	if err != nil {
		return domain.Session{}, &domain.ErrorFailedToInvalidateToken
	}

	profile, err := self.i.Persistence.FindProfileByUserID(token.User.ID)

	if err != nil {
		return domain.Session{}, &domain.ErrorProfileNotFound
	}

	derr := self.SetProfileRole(auditID, profile.ID, domain.RoleProfileStandard)

	if derr != nil {
		return domain.Session{}, &domain.ErrUnableToSetProfileRole
	}

	localizedDescription := self.i.Localize.GetLocalizedText(profile.Language, domain.TranslatableDefaultOrganization, domain.TranslatableArgumentMap{
		domain.TranslatableArgumentUserFullName: profile.Name,
	})

	organizationToCreate := domain.NewOrganization{
		Name:        profile.Name,
		Description: localizedDescription,
	}

	_, derr = self.CreateOrganization(auditID, profile.ID, organizationToCreate)

	if derr != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Session{}, derr
	}

	session, derr := self.createSession(auditID, "", token.User)
	if derr != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Session{}, derr
	}

	self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileMail, token.User.Email).Msg("Profile created")
	self.i.Trace.EndStep(auditID, auditStepID)
	return session, nil
}

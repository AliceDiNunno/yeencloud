package usecases

import (
	"time"

	"github.com/AliceDiNunno/yeencloud/src/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecases interface {
	CreateUser(auditID domain.AuditTraceID, user domain.NewUser, language string) (domain.User, *domain.ErrorDescription)

	GetUserByID(auditID domain.AuditTraceID, userID domain.UserID) (domain.User, *domain.ErrorDescription)
}

func (self UCs) newUserID() domain.UserID {
	return domain.UserID(uuid.New().String())
}

func (self UCs) CreateUser(auditID domain.AuditTraceID, newUser domain.NewUser, profileLanguage string) (domain.User, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, newUser.Secure())

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileMail, newUser.Email).Msg("Error hashing password")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.User{}, &domain.ErrorUnableToHashPassword
	}

	userToCreate := domain.User{
		ID:       self.newUserID(),
		Email:    newUser.Email,
		Password: string(hashedPassword),
	}

	user, err := self.i.Persistence.CreateUser(userToCreate)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileMail, newUser.Email).WithField(domain.LogFieldError, err).Msg("Error creating user")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.User{}, &domain.ErrorUnableToCreateUser
	}

	_, derr := self.createProfile(auditID, user.ID, newUser.Name, profileLanguage)

	if derr != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.User{}, derr
	}

	token := domain.Token{
		ID:        domain.TokenID(uuid.New().String()),
		User:      user,
		Token:     self.generateToken(),
		CreatedAt: time.Now().Unix(),
		ExpireAt:  time.Now().Add(time.Minute * 10).Unix(),
		Type:      domain.TokenTypeVerifyEmail,
	}

	_, err = self.i.Persistence.CreateToken(token)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileMail, newUser.Email).WithField(domain.LogFieldError, err).Msg("Error sending verification mail")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.User{}, &domain.ErrorUnableToGenerateToken
	}

	err = self.i.Mailer.SendVerificationMail(user.Email, token.Token)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileMail, newUser.Email).WithField(domain.LogFieldError, err).Msg("Error sending verification mail")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.User{}, &domain.ErrorUnableToSendVerificationMail
	}

	return user, nil
}

func (self UCs) GetUserByID(auditID domain.AuditTraceID, id domain.UserID) (domain.User, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip)
	user, err := self.i.Persistence.FindUserByID(id)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithLevel(domain.LogLevelError).WithField(domain.LogFieldError, err).WithField(domain.LogFieldUserID, id).Msg("Error finding user")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.User{}, &domain.ErrorUserNotFound
	}
	self.i.Trace.EndStep(auditID, auditStepID)
	return user, nil
}

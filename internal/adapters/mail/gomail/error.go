package gomail

import (
	"fmt"

	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
)

// MARK: - Objects
type MailerError struct {
	Msg string
	Key domain.Translatable
}

func (e *MailerError) Error() string {
	return fmt.Sprintf("mailer: %v", e.Msg)
}

func (e *MailerError) RawKey() domain.Translatable {
	return e.Key
}

func (e *MailerError) RestCode() int {
	return 502
}

// MARK: - Translatable
var (
	TranslatableUnableToSendVerificationMail = domain.Translatable{Key: "UnableToSendVerificationMail"}
)

// MARK: - Errors
var (
	ErrUnableToSendVerificationMail = MailerError{Msg: "unable to send verification mail", Key: TranslatableUnableToSendVerificationMail}
)

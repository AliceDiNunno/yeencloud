package interactor

type Mailer interface {
	SendVerificationMail(to string, token string) error
}

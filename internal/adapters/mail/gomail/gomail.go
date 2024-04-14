package gomail

import (
	"bytes"
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/AliceDiNunno/yeencloud/internal/core/domain/config"
	"github.com/davecgh/go-spew/spew"
	gomail "gopkg.in/mail.v2"
)

type Mailer struct {
	config config.MailConfig
}

type RegistrationValidation struct {
	Name     string
	Code     string
	IP       string
	Browser  string
	Location string

	ContactEmail string
}

func (m *Mailer) templateDirectory() string {
	if strings.HasPrefix(m.config.TemplateDirectory, "/") {
		return m.config.TemplateDirectory
	}

	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	return path.Join(wd, m.config.TemplateDirectory)
}

func (m *Mailer) executeTemplateFile(file string, data any) (string, error) {
	tmpl, err := template.New(file).ParseFiles(path.Join(m.templateDirectory(), file))
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (m *Mailer) sender() string {
	// TODO: get the name from somewhere else (domain?) or from the config (empty then set to a default)
	name := "YeenCloud"

	// TODO: get the environment from the config
	if os.Getenv("ENV") == "development" {
		name = name + "- Dev"
	}

	return name + " <" + m.config.From + ">"
}

func (m *Mailer) sendMail(to, subject, file string, data any) error {
	message := gomail.NewMessage()

	templateResult, err := m.executeTemplateFile(file, data)

	if err != nil {
		spew.Dump(err)
		return err
	}

	message.SetHeader("From", m.sender())
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", templateResult)

	d := gomail.NewDialer(m.config.Host, m.config.Port, m.config.From, m.config.Password)
	if err := d.DialAndSend(message); err != nil {
		return err
	}

	return nil
}

func (m *Mailer) SendVerificationMail(to string, token string) error {
	validation := RegistrationValidation{
		Name:     to,
		Code:     token,
		IP:       "85.232.120.20",
		Browser:  "Safari",
		Location: "Gif-sur-Yvette, France",

		ContactEmail: m.config.From,
	}
	var tmplFile = "verify_account.tpl.html"

	return m.sendMail(to, "Yeencloud - Verify your email", tmplFile, validation)
}

func NewMailer(config config.MailConfig) *Mailer {
	return &Mailer{
		config: config,
	}
}

package infrastructure

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(to, token, smtpHost, smtpUser, smtpPassword string, smtpPort int) error {
	verificationURL := fmt.Sprintf("http://localhost:8080/users/verify-email?token=%s", token)

	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Verify Your Email Address")
	m.SetBody("text/html", fmt.Sprintf(
		`<p>Please click the following link to verify your email address:</p>
        <a href="%s">Verify Email</a>`,
		verificationURL,
	))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendPassworkResetEmail(to, token, smtpHost, smtpUser, smtpPassword string, smtpPort int) error {
	verificationURL := fmt.Sprintf("http://localhost:8080/users/password-reset?token=%s", token)

	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "reset Your Password")
	m.SetBody("text/html", fmt.Sprintf(
		`<p>Please click the following link to reset Your Password"):</p>
        <a href="%s">Verify Email</a>`,
		verificationURL,
	))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

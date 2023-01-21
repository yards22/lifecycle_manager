package mailer

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type GoMail struct {
	mailSender gomail.SendCloser
	fromEmail  string
}

func NewGoMail(host string, port int, email string, password string, insecureSkipVerify bool) (*GoMail, error) {
	goMail := &GoMail{}
	var d = gomail.NewDialer(host, port, email, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: insecureSkipVerify}
	s, err := d.Dial()
	if err != nil {
		return nil, err
	}
	goMail.mailSender = s
	goMail.fromEmail = email
	return goMail, nil
}

func (gm *GoMail) Send(fromName, to, subject, body string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", fromName+"<"+gm.fromEmail+">")
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)
	err := gomail.Send(gm.mailSender, mail)
	if err != nil {
		return err
	}
	return nil
}

func (gm *GoMail) Close() error {
	return gm.mailSender.Close()
}

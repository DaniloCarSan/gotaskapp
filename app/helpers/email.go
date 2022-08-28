package helpers

import (
	"gotaskapp/app/config"

	mail "github.com/xhit/go-simple-mail/v2"
)

func creatServer() (*mail.SMTPClient, error) {

	server := mail.NewSMTPClient()
	server.Host = config.EMAIL_HOST
	server.Port = config.EMAIL_PORT
	server.Username = config.EMAIL_FROM
	server.Password = config.EMAIL_PASSWORD
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()

	if err != nil {
		return &mail.SMTPClient{}, err
	}

	return smtpClient, nil
}

func SendEmail(addTo []string, addCc []string, subject string, body string) error {

	smtpClient, err := creatServer()

	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom("Go TaskApp <" + config.EMAIL_FROM + ">")
	email.AddTo(addTo...)
	email.AddCc(addCc...)
	email.SetSubject(subject)
	email.SetBody(mail.TextHTML, body)

	err = email.Send(smtpClient)

	if err != nil {
		return err
	}

	return nil
}

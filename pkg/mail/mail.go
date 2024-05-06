package mail

import (
	"log"

	"github.com/DavidSie/notification-service/pkg/model"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Mailer struct {
	AppConfig model.AppConfig
}

func (m Mailer) ListenForMail() {
	go func() {
		for {
			msg := <-m.AppConfig.MailChannel
			err := m.sendMail(msg)
			if err != nil {
				log.Printf("Error while sending mail %v \n", err)
			}
		}

	}()

}

func (m Mailer) Send(emailRequest model.EmailRequest) error {
	return m.sendMail(emailRequest)
}

func (m Mailer) sendMail(emailRequest model.EmailRequest) error {
	server := mail.NewSMTPClient()
	server.Host = m.AppConfig.Stmp.Host
	server.Port = m.AppConfig.Stmp.Port
	server.Authentication = m.AppConfig.Stmp.Authentication
	server.Username = m.AppConfig.Stmp.Username
	server.Password = m.AppConfig.Stmp.Password
	server.TLSConfig = m.AppConfig.Stmp.TLSConfig
	server.Encryption = m.AppConfig.Stmp.Encryption
	server.KeepAlive = m.AppConfig.Stmp.KeepAlive
	server.ConnectTimeout = m.AppConfig.Stmp.ConnectTimeout
	server.SendTimeout = m.AppConfig.Stmp.SendTimeout

	client, err := server.Connect()

	if err != nil {
		return err
	}
	email := mail.NewMSG()
	email.SetFrom(emailRequest.Sender).AddTo(emailRequest.Recipients...).AddCc(emailRequest.CcRecipients...).AddBcc(emailRequest.BccRecipients...).SetSubject(emailRequest.Title)
	email.SetBody(emailRequest.MessageContentType, emailRequest.Message)
	err = email.Send(client)

	if err != nil {
		return err
	} else {
		log.Printf("Email to %s sent! \n", emailRequest.Recipients[0])
	}

	return nil
}

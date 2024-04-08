package mail

import (
	"fmt"

	"github.com/DavidSie/notification-service/pkg/model"
)

type Mailer struct{}

func (m Mailer) Send(emailRequest model.EmailRequest) error {

	fmt.Printf("Message send to %s title: %s \n ", emailRequest.Recipients, emailRequest.Title)
	return nil
	//auth smtp.PlainAuth()
}

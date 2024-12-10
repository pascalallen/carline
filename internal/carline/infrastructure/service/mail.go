package service

import (
	"fmt"
	mail2 "github.com/pascalallen/carline/internal/carline/domain/mail"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
)

type SendGridMailService struct{}

func NewSendGridMailService() mail2.Service {
	return &SendGridMailService{}
}

// TODO: refactor
func (s *SendGridMailService) Send(to string, subject string, body string) error {
	from := mail.NewEmail("Example User", "pascal.allen88@gmail.com")
	t := mail.NewEmail("Example User", to)
	message := mail.NewSingleEmail(from, subject, t, body, "")
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return nil
}

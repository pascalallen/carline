package mail

import (
	"github.com/sendgrid/sendgrid-go"
	"os"
)

func NewSendGridMailClient() *sendgrid.Client {
	return sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
}

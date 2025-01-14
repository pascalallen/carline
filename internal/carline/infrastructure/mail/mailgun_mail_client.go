package mail

import (
	"github.com/mailgun/mailgun-go/v4"
	"os"
)

func NewMailgunMailClient() *mailgun.MailgunImpl {
	return mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"))
}

package mail

type Service interface {
	Send(to string, subject string, body string) error
}

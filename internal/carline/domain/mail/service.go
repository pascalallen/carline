package mail

type Service interface {
	Send(from Sender, to Recipient, message Message) error // TODO: Determine if `from` is necessary as a param
}

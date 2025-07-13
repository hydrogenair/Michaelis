package mail

type Driver interface {
	Send(username string, to string, code string) error
}

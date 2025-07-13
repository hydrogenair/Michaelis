package mail

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"sync"
)

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

func NewMailer(host string, port int, username string, password string) *Mailer {
	dialer := gomail.NewDialer(host, port, username, password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{
				Dialer: dialer,
			},
		}
	})
	return internalMailer
}

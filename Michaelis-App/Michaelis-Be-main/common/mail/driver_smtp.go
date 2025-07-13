package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

type SMTP struct {
	Dialer *gomail.Dialer
}

func (smtp *SMTP) Send(username string, to string, code string) error {
	email := gomail.NewMessage()
	message := fmt.Sprintf("<h1> 欢迎使用聆析APP💗💗，您的验证码为 %s ，有效时间为5分钟。🥰</h1>", code)
	email.SetHeader("From", username)
	email.SetHeader("To", to)
	email.SetHeader("Subject", "聆析APP注册")
	email.SetBody("text/html", message)

	return smtp.Dialer.DialAndSend(email)
}

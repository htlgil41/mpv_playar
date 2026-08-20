package helpers

import (
	"fmt"
	"net/smtp"
	"playar/internal/libs"
)

func SendEmailServerExit(
	config *libs.ConfigApp,
	asunto string, suc string,
	body string,
) {

	if config.MailsSendersNotifier == nil {
		return
	}

	from := config.MailsSendersNotifier.From
	password := config.MailsSendersNotifier.Password_to
	to := []string{""}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := fmt.Sprintf("Hola te notifico que Playar API se ha detenido en la sucursal [%s]", suc)
	message := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return
	}
	return
}

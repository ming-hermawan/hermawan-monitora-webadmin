package hmonmail

import (
    "gopkg.in/gomail.v2"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
)

func Send(strTo string, subject string, message string) error {
    settings := dbo.GetSetting()
    smtpHost := settings.SmtpHost
    smtpPort := settings.SmtpPort
    senderName := settings.SenderName
    authEmail := settings.AuthEmail
    authPassword := settings.AuthPassword

    m := gomail.NewMessage()
    m.SetHeader("From", authEmail, senderName)
    m.SetHeader("To", strTo)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", message)

    d := gomail.NewDialer(smtpHost, smtpPort, authEmail, authPassword)

    return d.DialAndSend(m)
}

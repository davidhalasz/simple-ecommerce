package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

//go:embed templates
var emailTemplateFS embed.FS

func (app *application) SendMail(from, to, subject, tmpl string, data interface{}) error {
	templateToRender := fmt.Sprintf("templates/%s.html.tmpl", tmpl)

	t, err := template.New("email-html").ParseFS(emailTemplateFS, templateToRender)
	if err != nil {
		app.infoLog.Println("error 1")
		app.errorLog.Println(err)
		return err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", data); err != nil {
		app.infoLog.Println("error 2")
		app.errorLog.Println(err)
		return err
	}

	formattedMessage := tpl.String()

	templateToRender = fmt.Sprintf("templates/%s.plain.tmpl", tmpl)
	t, err = template.New("email-plain").ParseFS(emailTemplateFS, templateToRender)
	if err != nil {
		app.infoLog.Println("error 2")
		app.errorLog.Println(err)
		return err
	}

	if err = t.ExecuteTemplate(&tpl, "body", data); err != nil {
		app.infoLog.Println("error 3")
		app.errorLog.Println(err)
		return err
	}

	plainMessage := tpl.String()
	app.infoLog.Println(formattedMessage, plainMessage)

	// send the mail
	server := mail.NewSMTPClient()
	server.Host = app.config.smtp.host
	server.Port = app.config.smtp.port
	server.Username = app.config.smtp.username
	server.Password = app.config.smtp.password
	server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		app.infoLog.Println("error 4")
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(from).AddTo(to).SetSubject(subject)

	email.SetBody(mail.TextHTML, formattedMessage)
	email.AddAlternative(mail.TextPlain, plainMessage)

	err = email.Send(smtpClient)
	if err != nil {
		app.infoLog.Println("error 5")
		app.errorLog.Println(err)
		return err
	}

	app.infoLog.Println("send mail")

	return nil
}

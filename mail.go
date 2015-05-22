package main

import (
	"bytes"
	"net/smtp"
	"strconv"
	"strings"
	"text/template"
)

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

const MAIL_TEMPLATE = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}

Sincerely,

{{.From}}
`

func (m Mail) Send() error {
	t := template.Must(template.New("mail").Parse(MAIL_TEMPLATE))
	var b bytes.Buffer
	err := t.Execute(&b, m)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth(
		"",
		Username,
		Password,
		ServerAddr,
	)

	err = smtp.SendMail(
		ServerAddr+":"+strconv.Itoa(ServerPort),
		auth,
		From,
		strings.Split(To, ","),
		b.Bytes(),
	)

	return err
}

package email

import (
	"fmt"
	"net/smtp"
	"os"
)

type Client struct {
	host     string
	port     string
	username string
	password string
	from     string
}

func NewClient() *Client {
	return &Client{
		host:     os.Getenv("SMTP_HOST"),
		port:     os.Getenv("SMTP_PORT"),
		username: os.Getenv("SMTP_USERNAME"),
		password: os.Getenv("SMTP_PASSWORD"),
		from:     os.Getenv("SMTP_FROM"),
	}
}

func (c *Client) SendEmail(to []string, subject, body string) error {
	if c.host == "" {
		// If SMTP not configured, log but don't fail (useful for dev without Mailpit)
		fmt.Printf("SMTP not configured. Email to %v suppressed. Body: %s\n", to, body)
		return nil
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n%s\n%s", to[0], subject, mime, body))
	auth := smtp.PlainAuth("", c.username, c.password, c.host)
	addr := fmt.Sprintf("%s:%s", c.host, c.port)

	return smtp.SendMail(addr, auth, c.from, to, msg)
}

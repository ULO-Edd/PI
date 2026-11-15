package email

import (
	"log"
)

func (c *Client) TriggerRegistrationEmail(email, name, eventTitle string) {
	go func() {
		body := GetRegistrationTemplate(name, eventTitle)
		err := c.SendEmail([]string{email}, "Registration Confirmation: "+eventTitle, body)
		if err != nil {
			log.Printf("Failed to send registration email to %s: %v", email, err)
		}
	}()
}

func (c *Client) TriggerCancellationEmail(email, name, eventTitle string) {
	go func() {
		body := GetCancellationTemplate(name, eventTitle)
		err := c.SendEmail([]string{email}, "Event Cancelled: "+eventTitle, body)
		if err != nil {
			log.Printf("Failed to send cancellation email to %s: %v", email, err)
		}
	}()
}

func (c *Client) TriggerApprovalEmail(email, name string) {
	go func() {
		body := GetApprovalTemplate(name)
		err := c.SendEmail([]string{email}, "Account Approved", body)
		if err != nil {
			log.Printf("Failed to send approval email to %s: %v", email, err)
		}
	}()
}

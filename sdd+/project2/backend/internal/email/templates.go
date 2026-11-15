package email

import (
	"fmt"
)

func GetRegistrationTemplate(userName, eventTitle string) string {
	return fmt.Sprintf(`
		<h1>Registration Confirmed!</h1>
		<p>Hi %s,</p>
		<p>You have successfully registered for the event: <strong>%s</strong>.</p>
		<p>We look forward to seeing you there!</p>
	`, userName, eventTitle)
}

func GetCancellationTemplate(userName, eventTitle string) string {
	return fmt.Sprintf(`
		<h1>Event Cancelled</h1>
		<p>Hi %s,</p>
		<p>We regret to inform you that the event <strong>%s</strong> has been cancelled.</p>
		<p>We apologize for any inconvenience.</p>
	`, userName, eventTitle)
}

func GetApprovalTemplate(userName string) string {
	return fmt.Sprintf(`
		<h1>Account Approved</h1>
		<p>Hi %s,</p>
		<p>Your organizer account has been approved by an administrator. You can now publish events!</p>
	`, userName)
}

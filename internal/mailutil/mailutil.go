package mailutil

import(
	"fmt"
	"net/smtp"
)

type EmailConfig struct {
	SMTPHost string 
	SMTPPort string 
	Sender string
	Password string 
}

func SendResetTokenEmail (config *EmailConfig, token string, toEmail string) error{
	auth := smtp.PlainAuth("", config.Sender, config.Password,config.SMTPHost)

	subject := "Subject: TestTracker - Password Reset Token\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(`
		<h3>TestTracker Password Reset Request</h3>
		<p>You requested a password reset. Please use the following random string verification token to confirm your identity:</p>
		<p style="font-size: 18px; font-weight: bold; color: #4F46E5; letter-spacing: 1px;">%s</p>
		<p>This token will expire in 15 minutes. If you did not make this request, please ignore this email.</p>
	`, token)

	msg := []byte(subject + mime + body)
	addr := fmt.Sprintf("%s%s", config.SMTPHost,config.SMTPPort)

	err := smtp.SendMail(addr, auth, config.Sender, []string{toEmail}, msg)
	if err != nil {
		return fmt.Errorf("failed to send raw email payload: %w", err)
	}

	return nil


}
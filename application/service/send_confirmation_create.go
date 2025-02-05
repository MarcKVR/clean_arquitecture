package service

type Mailer interface {
	SendConfirmationEmail(email string) error
}

func SendConfirmationEmail(mail Mailer, email string) error {
	return mail.SendConfirmationEmail(email)
}

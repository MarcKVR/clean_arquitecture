package mailer

import "fmt"

// FakeMailer es una implementación de prueba para enviar emails.
type FakeMailer struct{}

// SendConfirmationEmail envía un email de confirmación.
// parámetros:
//
//	name: nombre del destinatario.
//	to: dirección de email del destinatario.
//	confirmationCode: código de confirmación que se enviará en el email.
func (f *FakeMailer) SendConfirmationEmail(name string, to string, confirmationCode string) error {
	// Simulación de envío de email
	fmt.Printf("Enviando email a %s (%s) con el código de confirmación: %s\n", name, to, confirmationCode)
	// Aquí se podría implementar la lógica real de envío de email.
	return nil
}

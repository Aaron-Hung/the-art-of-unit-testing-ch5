package ch5

type IEmailService interface {
	SendEmail(to string, body string, subject string)
}
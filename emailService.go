package ch5

type IEmailService interface {
	SendEmail(to string, body string, subject string)
}

type FakeEmailService struct {
	To string
	Body string
	Subject string
}

func (fe *FakeEmailService) SendEmail(to string, body string, subject string){
	fe.To = to
	fe.Body = body
	fe.Subject = subject
}
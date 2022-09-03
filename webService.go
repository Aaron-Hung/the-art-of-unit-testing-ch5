package ch5

type IWebService interface {
	LogError(message string)
}

type FakeWebService struct {
	lastError string
}

func (fw *FakeWebService) LogError(message string){
	fw.lastError = message;
}
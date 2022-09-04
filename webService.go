package ch5

import "errors"

type IWebService interface {
	LogError(message string) error
}

type FakeWebService struct {
	lastError string
	ShouldThrowError bool
}

func (fw *FakeWebService) LogError(message string) error {
	fw.lastError = message;
	if (fw.ShouldThrowError) {
		return errors.New("Fake Exception")
	}
	return nil;
}

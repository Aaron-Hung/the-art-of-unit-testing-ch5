package ch5

type IWebService interface {
	LogError(message string) error
}
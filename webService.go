package ch5

import "ch5/er"

type IWebService interface {
	LogError(message *er.ErrorInfo) error
}
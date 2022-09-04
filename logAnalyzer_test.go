package ch5

import (
	"ch5/mocks"
	"errors"
	"testing"
	"github.com/stretchr/testify/mock"
)

func Test_TooShortFileNameShouldSendEmail(t *testing.T) {
	mockWebService := mocks.NewIWebService(t)  // 建立假物件
	mockEmailService := mocks.NewIEmailService(t)  // 建立假物件
	la := NewLogAnalyzer(mockWebService, mockEmailService)
	mockWebService.On("LogError", mock.Anything).Return(errors.New("Fake Exception")).Once()
	
	expected_to := "support@going.cloud"
	expected_body := "Fake exception!"
	expected_subject := "Can't Log"
	mockEmailService.On("SendEmail", expected_to, expected_body, expected_subject).Once()

	tooShortFileName := "abc.txt"
	la.Analyze(tooShortFileName)
}

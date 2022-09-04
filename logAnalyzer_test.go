package ch5

import (
	"ch5/mocks"
	"ch5/er"
	"errors"
	"testing"
	"github.com/stretchr/testify/mock"
)

func Test_TooShortFileNameShouldSendEmail(t *testing.T) {
	mockWebService := mocks.NewIWebService(t)  // 建立假物件
	mockEmailService := mocks.NewIEmailService(t)  // 建立假物件
	la := NewLogAnalyzer(mockWebService, mockEmailService)
	// 驗證物件是帶著某些屬性的情況_方法1
	mockWebService.On("LogError", mock.MatchedBy(func(errorInfo *er.ErrorInfo) bool {
		return errorInfo.Filename == "abc.txt"
	})).Return(errors.New("Fake Exception")).Once()
	// 驗證物件是帶著某些屬性的情況_方法2
	// errorInfo := er.NewErrorInfo("File name too short", "abc.txt")
	// mockWebService.On("LogError", errorInfo).Return(errors.New("Fake Exception")).Once()
	
	expected_to := "support@going.cloud"
	expected_body := "Fake exception!"
	expected_subject := "Can't Log"
	mockEmailService.On("SendEmail", expected_to, expected_body, expected_subject).Once()

	tooShortFileName := "abc.txt"
	la.Analyze(tooShortFileName)
}

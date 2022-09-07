package ch5

import (
	"ch5/mocks"
	"testing"
)

func Test_TooShortFileNameShouldLogError(t *testing.T) {
	mockWebService := mocks.NewIWebService(t)  // 建立假物件
	la := NewLogAnalyzer(mockWebService)
	expected := "File name too short: " + "abc.txt"
	mockWebService.On("LogError", expected).Once()

	tooShortFileName := "abc.txt"
	la.Analyze(tooShortFileName)
}
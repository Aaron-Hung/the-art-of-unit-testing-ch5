package ch5

import (
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/golang/mock/gomock"
)

func Test_TooShortFileNameShouldLogError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockIWebService := NewMockIWebService(ctl)
	expected := "File name too short: " + "abc.txt"
	gomock.InOrder(
		mockIWebService.EXPECT().LogError(expected),
	)
	la := NewLogAnalyzer(mockIWebService)

	tooShortFileName := "abc.txt"
	la.Analyze(tooShortFileName)
}
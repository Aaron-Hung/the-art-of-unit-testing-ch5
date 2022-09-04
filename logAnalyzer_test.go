package ch5

import (
	"testing"
	"errors"
	"github.com/golang/mock/gomock"
)

func Test_TooShortFileNameShouldSendEmail(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockIWebService := NewMockIWebService(ctl)
	mockIEmailService := NewMockIEmailService(ctl)

	expected_to := "support@going.cloud"
	expected_body := "Fake exception!"
	expected_subject := "Can't Log"

	gomock.InOrder(
		mockIWebService.EXPECT().LogError(gomock.Any()).Return(errors.New("Fake Exception")).
		Times(1),
		// MinTimes(1).MaxTimes(10),

		mockIEmailService.EXPECT().SendEmail(expected_to, expected_body, expected_subject).
		Times(1),
		// AnyTimes(),
	)

	la := NewLogAnalyzer(mockIWebService, mockIEmailService)

	tooShortFileName := "abc.txt"
	la.Analyze(tooShortFileName)
}

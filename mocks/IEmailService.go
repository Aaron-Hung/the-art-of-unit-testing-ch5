// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IEmailService is an autogenerated mock type for the IEmailService type
type IEmailService struct {
	mock.Mock
}

// SendEmail provides a mock function with given fields: to, body, subject
func (_m *IEmailService) SendEmail(to string, body string, subject string) {
	_m.Called(to, body, subject)
}

type mockConstructorTestingTNewIEmailService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIEmailService creates a new instance of IEmailService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIEmailService(t mockConstructorTestingTNewIEmailService) *IEmailService {
	mock := &IEmailService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
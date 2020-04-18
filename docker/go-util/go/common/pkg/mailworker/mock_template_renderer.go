// Code generated by mockery v1.0.0. DO NOT EDIT.
package mailworker

import (
	mock "github.com/stretchr/testify/mock"
)

// MockTemplateRenderer is an autogenerated mock type for the TemplateRenderer type
type MockTemplateRenderer struct {
	mock.Mock
}

// RenderFileWithKeyValuePairs provides a mock function with given fields: filename, templateData
func (_m *MockTemplateRenderer) RenderFileWithKeyValuePairs(filename string, templateData map[string]string) (string, error) {
	ret := _m.Called(filename, templateData)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, map[string]string) string); ok {
		r0 = rf(filename, templateData)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]string) error); ok {
		r1 = rf(filename, templateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

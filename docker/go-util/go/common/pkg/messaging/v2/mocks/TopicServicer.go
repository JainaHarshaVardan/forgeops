// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	pubsub "cloud.google.com/go/pubsub"
)

// TopicServicer is an autogenerated mock type for the TopicServicer type
type TopicServicer struct {
	mock.Mock
}

// EnsureExists provides a mock function with given fields: _a0
func (_m *TopicServicer) EnsureExists(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: ctx, msg
func (_m *TopicServicer) Publish(ctx context.Context, msg *pubsub.Message) *pubsub.PublishResult {
	ret := _m.Called(ctx, msg)

	var r0 *pubsub.PublishResult
	if rf, ok := ret.Get(0).(func(context.Context, *pubsub.Message) *pubsub.PublishResult); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pubsub.PublishResult)
		}
	}

	return r0
}

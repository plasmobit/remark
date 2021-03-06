// Code generated by mockery v1.0.0. DO NOT EDIT.
package image

import context "context"
import io "io"
import mock "github.com/stretchr/testify/mock"
import time "time"

// MockStore is an autogenerated mock type for the Store type
type MockStore struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: ctx, ttl
func (_m *MockStore) Cleanup(ctx context.Context, ttl time.Duration) error {
	ret := _m.Called(ctx, ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration) error); ok {
		r0 = rf(ctx, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Commit provides a mock function with given fields: id
func (_m *MockStore) Commit(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Load provides a mock function with given fields: id
func (_m *MockStore) Load(id string) (io.ReadCloser, int64, error) {
	ret := _m.Called(id)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string) io.ReadCloser); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(string) int64); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Save provides a mock function with given fields: fileName, userID, r
func (_m *MockStore) Save(fileName string, userID string, r io.Reader) (string, error) {
	ret := _m.Called(fileName, userID, r)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) string); ok {
		r0 = rf(fileName, userID, r)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, io.Reader) error); ok {
		r1 = rf(fileName, userID, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SizeLimit provides a mock function with given fields:
func (_m *MockStore) SizeLimit() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

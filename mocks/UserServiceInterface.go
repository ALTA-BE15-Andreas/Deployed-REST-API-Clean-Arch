// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	users "clean-arch/features/users"

	mock "github.com/stretchr/testify/mock"
)

// UserServiceMock is an autogenerated mock type for the UserServiceInterface type
type UserServiceMock struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *UserServiceMock) Create(input users.UserEntity) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.UserEntity) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *UserServiceMock) GetAll() ([]users.UserEntity, error) {
	ret := _m.Called()

	var r0 []users.UserEntity
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]users.UserEntity, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []users.UserEntity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.UserEntity)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: userId
func (_m *UserServiceMock) GetUser(userId uint) (users.UserEntity, error) {
	ret := _m.Called(userId)

	var r0 users.UserEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (users.UserEntity, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(uint) users.UserEntity); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(users.UserEntity)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modify provides a mock function with given fields: userId, input
func (_m *UserServiceMock) Modify(userId uint, input users.UserEntity) error {
	ret := _m.Called(userId, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.UserEntity) error); ok {
		r0 = rf(userId, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remove provides a mock function with given fields: userId
func (_m *UserServiceMock) Remove(userId uint) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserServiceMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserServiceMock creates a new instance of UserServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserServiceMock(t mockConstructorTestingTNewUserServiceMock) *UserServiceMock {
	mock := &UserServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
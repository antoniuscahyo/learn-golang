// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/entity"
	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// CreateAsGuest provides a mock function with given fields: ctx, createUserReq
func (_m *UserUsecase) CreateAsGuest(ctx context.Context, createUserReq *entity.CreateUserReq) (*entity.CreateUserResponse, error) {
	ret := _m.Called(ctx, createUserReq)

	var r0 *entity.CreateUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *entity.CreateUserReq) *entity.CreateUserResponse); ok {
		r0 = rf(ctx, createUserReq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.CreateUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.CreateUserReq) error); ok {
		r1 = rf(ctx, createUserReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyByEmailAndPassword provides a mock function with given fields: ctx, req
func (_m *UserUsecase) VerifyByEmailAndPassword(ctx context.Context, req *entity.LoginReq) (*entity.LoginResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *entity.LoginResponse
	if rf, ok := ret.Get(0).(func(context.Context, *entity.LoginReq) *entity.LoginResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.LoginResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.LoginReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecase(t mockConstructorTestingTNewUserUsecase) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

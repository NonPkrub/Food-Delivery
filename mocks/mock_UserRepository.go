package mocks

import (
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

// func (_m *UserRepository) SignUp(u *domain.UserSignUpForm) (*domain.UserReply, error) {
// 	ret := _m.Called(t)

// 	var r0 error
// 	if rf, ok := ret.Get(0).(func(*domain.UserSignUpForm) (*domain.UserReply, error)); ok {
// 		r0 = rt(t)
// 	} else {
// 		r0 = ret.Error(0)
// 	}

// 	return r0
// }

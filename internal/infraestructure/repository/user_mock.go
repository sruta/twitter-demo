package repository

import (
	"github.com/stretchr/testify/mock"
	"twitter-demo/internal/domain"
	"twitter-demo/pkg"
)

type UserMock struct {
	mock.Mock
}

func NewUserMock() *UserMock {
	return &UserMock{}
}

func (u *UserMock) Select() ([]domain.User, pkg.Error) {
	args := u.Called()
	arg0 := args.Get(0).([]domain.User)
	if args.Get(1) != nil {
		return arg0, args.Get(1).(pkg.Error)
	}
	return arg0, nil
}

func (u *UserMock) SelectByID(id int64) (domain.User, pkg.Error) {
	args := u.Called(id)
	arg0 := args.Get(0).(domain.User)
	if args.Get(1) != nil {
		return arg0, args.Get(1).(pkg.Error)
	}
	return arg0, nil
}

func (u *UserMock) SelectByEmail(email string) (domain.User, pkg.Error) {
	args := u.Called(email)
	arg0 := args.Get(0).(domain.User)
	if args.Get(1) != nil {
		return arg0, args.Get(1).(pkg.Error)
	}
	return arg0, nil
}

func (u *UserMock) SelectByUsername(username string) (domain.User, pkg.Error) {
	args := u.Called(username)
	arg0 := args.Get(0).(domain.User)
	if args.Get(1) != nil {
		return arg0, args.Get(1).(pkg.Error)
	}
	return arg0, nil
}

func (u *UserMock) Insert(user domain.User) (domain.User, pkg.Error) {
	args := u.Called(user)
	arg0 := args.Get(0).(domain.User)
	if args.Get(1) != nil {
		return arg0, args.Get(1).(pkg.Error)
	}
	return arg0, nil
}

func (u *UserMock) Update(user domain.User) (domain.User, pkg.Error) {
	args := u.Called(user)
	arg0 := args.Get(0).(domain.User)
	if args.Get(1) != nil {
		return arg0, args.Get(1).(pkg.Error)
	}
	return arg0, nil
}

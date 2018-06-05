package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/t-tashiro/yakiniku/account"
)

// AccountRepository mock account repository
type AccountRepository struct {
	mock.Mock
}

// List list mock accounts
func (_m *AccountRepository) List() ([]*account.Account, error) {
	ret := _m.Called()

	var r0 []*account.Account
	if rf, ok := ret.Get(0).(func() []*account.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*account.Account)

		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

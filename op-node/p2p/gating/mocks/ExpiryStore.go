// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	net "net"

	mock "github.com/stretchr/testify/mock"

	peer "github.com/libp2p/go-libp2p/core/peer"

	time "time"
)

// ExpiryStore is an autogenerated mock type for the ExpiryStore type
type ExpiryStore struct {
	mock.Mock
}

type ExpiryStore_Expecter struct {
	mock *mock.Mock
}

func (_m *ExpiryStore) EXPECT() *ExpiryStore_Expecter {
	return &ExpiryStore_Expecter{mock: &_m.Mock}
}

// GetIPBanExpiration provides a mock function with given fields: ip
func (_m *ExpiryStore) GetIPBanExpiration(ip net.IP) (time.Time, error) {
	ret := _m.Called(ip)

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(net.IP) (time.Time, error)); ok {
		return rf(ip)
	}
	if rf, ok := ret.Get(0).(func(net.IP) time.Time); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(net.IP) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExpiryStore_GetIPBanExpiration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIPBanExpiration'
type ExpiryStore_GetIPBanExpiration_Call struct {
	*mock.Call
}

// GetIPBanExpiration is a helper method to define mock.On call
//   - ip net.IP
func (_e *ExpiryStore_Expecter) GetIPBanExpiration(ip interface{}) *ExpiryStore_GetIPBanExpiration_Call {
	return &ExpiryStore_GetIPBanExpiration_Call{Call: _e.mock.On("GetIPBanExpiration", ip)}
}

func (_c *ExpiryStore_GetIPBanExpiration_Call) Run(run func(ip net.IP)) *ExpiryStore_GetIPBanExpiration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.IP))
	})
	return _c
}

func (_c *ExpiryStore_GetIPBanExpiration_Call) Return(_a0 time.Time, _a1 error) *ExpiryStore_GetIPBanExpiration_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExpiryStore_GetIPBanExpiration_Call) RunAndReturn(run func(net.IP) (time.Time, error)) *ExpiryStore_GetIPBanExpiration_Call {
	_c.Call.Return(run)
	return _c
}

// GetPeerBanExpiration provides a mock function with given fields: id
func (_m *ExpiryStore) GetPeerBanExpiration(id peer.ID) (time.Time, error) {
	ret := _m.Called(id)

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(peer.ID) (time.Time, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) time.Time); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(peer.ID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExpiryStore_GetPeerBanExpiration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPeerBanExpiration'
type ExpiryStore_GetPeerBanExpiration_Call struct {
	*mock.Call
}

// GetPeerBanExpiration is a helper method to define mock.On call
//   - id peer.ID
func (_e *ExpiryStore_Expecter) GetPeerBanExpiration(id interface{}) *ExpiryStore_GetPeerBanExpiration_Call {
	return &ExpiryStore_GetPeerBanExpiration_Call{Call: _e.mock.On("GetPeerBanExpiration", id)}
}

func (_c *ExpiryStore_GetPeerBanExpiration_Call) Run(run func(id peer.ID)) *ExpiryStore_GetPeerBanExpiration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(peer.ID))
	})
	return _c
}

func (_c *ExpiryStore_GetPeerBanExpiration_Call) Return(_a0 time.Time, _a1 error) *ExpiryStore_GetPeerBanExpiration_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExpiryStore_GetPeerBanExpiration_Call) RunAndReturn(run func(peer.ID) (time.Time, error)) *ExpiryStore_GetPeerBanExpiration_Call {
	_c.Call.Return(run)
	return _c
}

// SetIPBanExpiration provides a mock function with given fields: ip, expiry
func (_m *ExpiryStore) SetIPBanExpiration(ip net.IP, expiry time.Time) error {
	ret := _m.Called(ip, expiry)

	var r0 error
	if rf, ok := ret.Get(0).(func(net.IP, time.Time) error); ok {
		r0 = rf(ip, expiry)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExpiryStore_SetIPBanExpiration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetIPBanExpiration'
type ExpiryStore_SetIPBanExpiration_Call struct {
	*mock.Call
}

// SetIPBanExpiration is a helper method to define mock.On call
//   - ip net.IP
//   - expiry time.Time
func (_e *ExpiryStore_Expecter) SetIPBanExpiration(ip interface{}, expiry interface{}) *ExpiryStore_SetIPBanExpiration_Call {
	return &ExpiryStore_SetIPBanExpiration_Call{Call: _e.mock.On("SetIPBanExpiration", ip, expiry)}
}

func (_c *ExpiryStore_SetIPBanExpiration_Call) Run(run func(ip net.IP, expiry time.Time)) *ExpiryStore_SetIPBanExpiration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.IP), args[1].(time.Time))
	})
	return _c
}

func (_c *ExpiryStore_SetIPBanExpiration_Call) Return(_a0 error) *ExpiryStore_SetIPBanExpiration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ExpiryStore_SetIPBanExpiration_Call) RunAndReturn(run func(net.IP, time.Time) error) *ExpiryStore_SetIPBanExpiration_Call {
	_c.Call.Return(run)
	return _c
}

// SetPeerBanExpiration provides a mock function with given fields: id, expiry
func (_m *ExpiryStore) SetPeerBanExpiration(id peer.ID, expiry time.Time) error {
	ret := _m.Called(id, expiry)

	var r0 error
	if rf, ok := ret.Get(0).(func(peer.ID, time.Time) error); ok {
		r0 = rf(id, expiry)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExpiryStore_SetPeerBanExpiration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPeerBanExpiration'
type ExpiryStore_SetPeerBanExpiration_Call struct {
	*mock.Call
}

// SetPeerBanExpiration is a helper method to define mock.On call
//   - id peer.ID
//   - expiry time.Time
func (_e *ExpiryStore_Expecter) SetPeerBanExpiration(id interface{}, expiry interface{}) *ExpiryStore_SetPeerBanExpiration_Call {
	return &ExpiryStore_SetPeerBanExpiration_Call{Call: _e.mock.On("SetPeerBanExpiration", id, expiry)}
}

func (_c *ExpiryStore_SetPeerBanExpiration_Call) Run(run func(id peer.ID, expiry time.Time)) *ExpiryStore_SetPeerBanExpiration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(peer.ID), args[1].(time.Time))
	})
	return _c
}

func (_c *ExpiryStore_SetPeerBanExpiration_Call) Return(_a0 error) *ExpiryStore_SetPeerBanExpiration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ExpiryStore_SetPeerBanExpiration_Call) RunAndReturn(run func(peer.ID, time.Time) error) *ExpiryStore_SetPeerBanExpiration_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewExpiryStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewExpiryStore creates a new instance of ExpiryStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExpiryStore(t mockConstructorTestingTNewExpiryStore) *ExpiryStore {
	mock := &ExpiryStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

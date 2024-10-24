// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/updevru/go-micro-kit-example/internal/domain"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// DeleteLog provides a mock function with given fields: item
func (_m *Client) DeleteLog(item *domain.ItemStore) {
	_m.Called(item)
}

// Client_DeleteLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteLog'
type Client_DeleteLog_Call struct {
	*mock.Call
}

// DeleteLog is a helper method to define mock.On call
//   - item *domain.ItemStore
func (_e *Client_Expecter) DeleteLog(item interface{}) *Client_DeleteLog_Call {
	return &Client_DeleteLog_Call{Call: _e.mock.On("DeleteLog", item)}
}

func (_c *Client_DeleteLog_Call) Run(run func(item *domain.ItemStore)) *Client_DeleteLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.ItemStore))
	})
	return _c
}

func (_c *Client_DeleteLog_Call) Return() *Client_DeleteLog_Call {
	_c.Call.Return()
	return _c
}

func (_c *Client_DeleteLog_Call) RunAndReturn(run func(*domain.ItemStore)) *Client_DeleteLog_Call {
	_c.Call.Return(run)
	return _c
}

// SaveLog provides a mock function with given fields: item
func (_m *Client) SaveLog(item *domain.ItemStore) {
	_m.Called(item)
}

// Client_SaveLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveLog'
type Client_SaveLog_Call struct {
	*mock.Call
}

// SaveLog is a helper method to define mock.On call
//   - item *domain.ItemStore
func (_e *Client_Expecter) SaveLog(item interface{}) *Client_SaveLog_Call {
	return &Client_SaveLog_Call{Call: _e.mock.On("SaveLog", item)}
}

func (_c *Client_SaveLog_Call) Run(run func(item *domain.ItemStore)) *Client_SaveLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.ItemStore))
	})
	return _c
}

func (_c *Client_SaveLog_Call) Return() *Client_SaveLog_Call {
	_c.Call.Return()
	return _c
}

func (_c *Client_SaveLog_Call) RunAndReturn(run func(*domain.ItemStore)) *Client_SaveLog_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

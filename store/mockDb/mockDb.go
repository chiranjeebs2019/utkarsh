// Code generated by mockery v1.0.0. DO NOT EDIT.

package mockDb

import mock "github.com/stretchr/testify/mock"

// DbStore is an autogenerated mock type for the DbStore type
type MockDb struct {
	mock.Mock
}

// Insert provides a mock function with given fields: data
func (_m *MockDb) Insert(data interface{}) bool {
	ret := _m.Called(data)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
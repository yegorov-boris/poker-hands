// Code generated by mockery v1.0.0. DO NOT EDIT.
package main

import mock "github.com/stretchr/testify/mock"

// mockSorterByValue is an autogenerated mock type for the sorterByValue type
type mockSorterByValue struct {
	mock.Mock
}

// sortByValue provides a mock function with given fields: _a0
func (_m *mockSorterByValue) sortByValue(_a0 Hand) Hand {
	ret := _m.Called(_a0)

	var r0 Hand
	if rf, ok := ret.Get(0).(func(Hand) Hand); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Hand)
		}
	}

	return r0
}
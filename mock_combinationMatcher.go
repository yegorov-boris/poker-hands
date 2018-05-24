// Code generated by mockery v1.0.0. DO NOT EDIT.
package main

import mock "github.com/stretchr/testify/mock"

// mockCombinationMatcher is an autogenerated mock type for the combinationMatcher type
type mockCombinationMatcher struct {
	mock.Mock
}

// isFlush provides a mock function with given fields: hand
func (_m *mockCombinationMatcher) isFlush(hand Hand) (bool, Hand) {
	ret := _m.Called(hand)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Hand) bool); ok {
		r0 = rf(hand)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 Hand
	if rf, ok := ret.Get(1).(func(Hand) Hand); ok {
		r1 = rf(hand)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(Hand)
		}
	}

	return r0, r1
}

// isFourKind provides a mock function with given fields: hand
func (_m *mockCombinationMatcher) isFourKind(hand Hand) (bool, Hand) {
	ret := _m.Called(hand)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Hand) bool); ok {
		r0 = rf(hand)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 Hand
	if rf, ok := ret.Get(1).(func(Hand) Hand); ok {
		r1 = rf(hand)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(Hand)
		}
	}

	return r0, r1
}

// isFullHouse provides a mock function with given fields: hand
func (_m *mockCombinationMatcher) isFullHouse(hand Hand) (bool, Hand) {
	ret := _m.Called(hand)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Hand) bool); ok {
		r0 = rf(hand)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 Hand
	if rf, ok := ret.Get(1).(func(Hand) Hand); ok {
		r1 = rf(hand)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(Hand)
		}
	}

	return r0, r1
}

// isOnePair provides a mock function with given fields: hand
func (_m *mockCombinationMatcher) isOnePair(hand Hand) (bool, Hand) {
	ret := _m.Called(hand)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Hand) bool); ok {
		r0 = rf(hand)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 Hand
	if rf, ok := ret.Get(1).(func(Hand) Hand); ok {
		r1 = rf(hand)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(Hand)
		}
	}

	return r0, r1
}

// isRoyalFlush provides a mock function with given fields: hand
func (_m *mockCombinationMatcher) isRoyalFlush(hand Hand) (bool, Hand) {
	ret := _m.Called(hand)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Hand) bool); ok {
		r0 = rf(hand)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 Hand
	if rf, ok := ret.Get(1).(func(Hand) Hand); ok {
		r1 = rf(hand)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(Hand)
		}
	}

	return r0, r1
}

// isStraight provides a mock function with given fields: hand
func (_m *mockCombinationMatcher) isStraight(hand Hand) (bool, Hand) {
	ret := _m.Called(hand)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Hand) bool); ok {
		r0 = rf(hand)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 Hand
	if rf, ok := ret.Get(1).(func(Hand) Hand); ok {
		r1 = rf(hand)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(Hand)
		}
	}

	return r0, r1
}

// isStraightFlush provides a mock function with given fields: hand
func (_m *mockCombinationMatcher) isStraightFlush(hand Hand) (bool, Hand) {
	ret := _m.Called(hand)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Hand) bool); ok {
		r0 = rf(hand)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 Hand
	if rf, ok := ret.Get(1).(func(Hand) Hand); ok {
		r1 = rf(hand)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(Hand)
		}
	}

	return r0, r1
}

// isThreeKind provides a mock function with given fields: hand
func (_m *mockCombinationMatcher) isThreeKind(hand Hand) (bool, Hand) {
	ret := _m.Called(hand)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Hand) bool); ok {
		r0 = rf(hand)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 Hand
	if rf, ok := ret.Get(1).(func(Hand) Hand); ok {
		r1 = rf(hand)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(Hand)
		}
	}

	return r0, r1
}

// isTwoPairs provides a mock function with given fields: hand
func (_m *mockCombinationMatcher) isTwoPairs(hand Hand) (bool, Hand) {
	ret := _m.Called(hand)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Hand) bool); ok {
		r0 = rf(hand)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 Hand
	if rf, ok := ret.Get(1).(func(Hand) Hand); ok {
		r1 = rf(hand)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(Hand)
		}
	}

	return r0, r1
}
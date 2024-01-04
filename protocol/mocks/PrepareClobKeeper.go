// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	clobtypes "github.com/dydxprotocol/v4-chain/protocol/x/clob/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// PrepareClobKeeper is an autogenerated mock type for the PrepareClobKeeper type
type PrepareClobKeeper struct {
	mock.Mock
}

// GetOperations provides a mock function with given fields: ctx
func (_m *PrepareClobKeeper) GetOperations(ctx types.Context) *clobtypes.MsgProposedOperations {
	ret := _m.Called(ctx)

	var r0 *clobtypes.MsgProposedOperations
	if rf, ok := ret.Get(0).(func(types.Context) *clobtypes.MsgProposedOperations); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clobtypes.MsgProposedOperations)
		}
	}

	return r0
}

type mockConstructorTestingTNewPrepareClobKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewPrepareClobKeeper creates a new instance of PrepareClobKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPrepareClobKeeper(t mockConstructorTestingTNewPrepareClobKeeper) *PrepareClobKeeper {
	mock := &PrepareClobKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

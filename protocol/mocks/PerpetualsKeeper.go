// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	perpetualstypes "github.com/dydxprotocol/v4-chain/protocol/x/perpetuals/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// PerpetualsKeeper is an autogenerated mock type for the PerpetualsKeeper type
type PerpetualsKeeper struct {
	mock.Mock
}

// AddPremiumVotes provides a mock function with given fields: ctx, votes
func (_m *PerpetualsKeeper) AddPremiumVotes(ctx types.Context, votes []perpetualstypes.FundingPremium) error {
	ret := _m.Called(ctx, votes)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, []perpetualstypes.FundingPremium) error); ok {
		r0 = rf(ctx, votes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePerpetual provides a mock function with given fields: ctx, id, ticker, marketId, atomicResolution, defaultFundingPpm, liquidityTier
func (_m *PerpetualsKeeper) CreatePerpetual(ctx types.Context, id uint32, ticker string, marketId uint32, atomicResolution int32, defaultFundingPpm int32, liquidityTier uint32) (perpetualstypes.Perpetual, error) {
	ret := _m.Called(ctx, id, ticker, marketId, atomicResolution, defaultFundingPpm, liquidityTier)

	var r0 perpetualstypes.Perpetual
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint32, string, uint32, int32, int32, uint32) (perpetualstypes.Perpetual, error)); ok {
		return rf(ctx, id, ticker, marketId, atomicResolution, defaultFundingPpm, liquidityTier)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint32, string, uint32, int32, int32, uint32) perpetualstypes.Perpetual); ok {
		r0 = rf(ctx, id, ticker, marketId, atomicResolution, defaultFundingPpm, liquidityTier)
	} else {
		r0 = ret.Get(0).(perpetualstypes.Perpetual)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint32, string, uint32, int32, int32, uint32) error); ok {
		r1 = rf(ctx, id, ticker, marketId, atomicResolution, defaultFundingPpm, liquidityTier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddPremiumVotes provides a mock function with given fields: ctx
func (_m *PerpetualsKeeper) GetAddPremiumVotes(ctx types.Context) *perpetualstypes.MsgAddPremiumVotes {
	ret := _m.Called(ctx)

	var r0 *perpetualstypes.MsgAddPremiumVotes
	if rf, ok := ret.Get(0).(func(types.Context) *perpetualstypes.MsgAddPremiumVotes); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*perpetualstypes.MsgAddPremiumVotes)
		}
	}

	return r0
}

// GetMarginRequirements provides a mock function with given fields: ctx, id, bigQuantums
func (_m *PerpetualsKeeper) GetMarginRequirements(ctx types.Context, id uint32, bigQuantums *big.Int) (*big.Int, *big.Int, error) {
	ret := _m.Called(ctx, id, bigQuantums)

	var r0 *big.Int
	var r1 *big.Int
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Context, uint32, *big.Int) (*big.Int, *big.Int, error)); ok {
		return rf(ctx, id, bigQuantums)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint32, *big.Int) *big.Int); ok {
		r0 = rf(ctx, id, bigQuantums)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint32, *big.Int) *big.Int); ok {
		r1 = rf(ctx, id, bigQuantums)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	if rf, ok := ret.Get(2).(func(types.Context, uint32, *big.Int) error); ok {
		r2 = rf(ctx, id, bigQuantums)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetNetCollateral provides a mock function with given fields: ctx, id, bigQuantums
func (_m *PerpetualsKeeper) GetNetCollateral(ctx types.Context, id uint32, bigQuantums *big.Int) (*big.Int, error) {
	ret := _m.Called(ctx, id, bigQuantums)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint32, *big.Int) (*big.Int, error)); ok {
		return rf(ctx, id, bigQuantums)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint32, *big.Int) *big.Int); ok {
		r0 = rf(ctx, id, bigQuantums)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint32, *big.Int) error); ok {
		r1 = rf(ctx, id, bigQuantums)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNetNotional provides a mock function with given fields: ctx, id, bigQuantums
func (_m *PerpetualsKeeper) GetNetNotional(ctx types.Context, id uint32, bigQuantums *big.Int) (*big.Int, error) {
	ret := _m.Called(ctx, id, bigQuantums)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint32, *big.Int) (*big.Int, error)); ok {
		return rf(ctx, id, bigQuantums)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint32, *big.Int) *big.Int); ok {
		r0 = rf(ctx, id, bigQuantums)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint32, *big.Int) error); ok {
		r1 = rf(ctx, id, bigQuantums)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotionalInBaseQuantums provides a mock function with given fields: ctx, id, bigQuoteQuantums
func (_m *PerpetualsKeeper) GetNotionalInBaseQuantums(ctx types.Context, id uint32, bigQuoteQuantums *big.Int) (*big.Int, error) {
	ret := _m.Called(ctx, id, bigQuoteQuantums)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint32, *big.Int) (*big.Int, error)); ok {
		return rf(ctx, id, bigQuoteQuantums)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint32, *big.Int) *big.Int); ok {
		r0 = rf(ctx, id, bigQuoteQuantums)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint32, *big.Int) error); ok {
		r1 = rf(ctx, id, bigQuoteQuantums)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasAuthority provides a mock function with given fields: authority
func (_m *PerpetualsKeeper) HasAuthority(authority string) bool {
	ret := _m.Called(authority)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(authority)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MaybeProcessNewFundingSampleEpoch provides a mock function with given fields: ctx
func (_m *PerpetualsKeeper) MaybeProcessNewFundingSampleEpoch(ctx types.Context) {
	_m.Called(ctx)
}

// MaybeProcessNewFundingTickEpoch provides a mock function with given fields: ctx
func (_m *PerpetualsKeeper) MaybeProcessNewFundingTickEpoch(ctx types.Context) {
	_m.Called(ctx)
}

// ModifyPerpetual provides a mock function with given fields: ctx, id, ticker, marketId, defaultFundingPpm, liquidityTier
func (_m *PerpetualsKeeper) ModifyPerpetual(ctx types.Context, id uint32, ticker string, marketId uint32, defaultFundingPpm int32, liquidityTier uint32) (perpetualstypes.Perpetual, error) {
	ret := _m.Called(ctx, id, ticker, marketId, defaultFundingPpm, liquidityTier)

	var r0 perpetualstypes.Perpetual
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint32, string, uint32, int32, uint32) (perpetualstypes.Perpetual, error)); ok {
		return rf(ctx, id, ticker, marketId, defaultFundingPpm, liquidityTier)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint32, string, uint32, int32, uint32) perpetualstypes.Perpetual); ok {
		r0 = rf(ctx, id, ticker, marketId, defaultFundingPpm, liquidityTier)
	} else {
		r0 = ret.Get(0).(perpetualstypes.Perpetual)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint32, string, uint32, int32, uint32) error); ok {
		r1 = rf(ctx, id, ticker, marketId, defaultFundingPpm, liquidityTier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PerformStatefulPremiumVotesValidation provides a mock function with given fields: ctx, msg
func (_m *PerpetualsKeeper) PerformStatefulPremiumVotesValidation(ctx types.Context, msg *perpetualstypes.MsgAddPremiumVotes) error {
	ret := _m.Called(ctx, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualstypes.MsgAddPremiumVotes) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLiquidityTier provides a mock function with given fields: ctx, id, name, initialMarginPpm, maintenanceFractionPpm, impactNotional
func (_m *PerpetualsKeeper) SetLiquidityTier(ctx types.Context, id uint32, name string, initialMarginPpm uint32, maintenanceFractionPpm uint32, impactNotional uint64) (perpetualstypes.LiquidityTier, error) {
	ret := _m.Called(ctx, id, name, initialMarginPpm, maintenanceFractionPpm, impactNotional)

	var r0 perpetualstypes.LiquidityTier
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint32, string, uint32, uint32, uint64) (perpetualstypes.LiquidityTier, error)); ok {
		return rf(ctx, id, name, initialMarginPpm, maintenanceFractionPpm, impactNotional)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint32, string, uint32, uint32, uint64) perpetualstypes.LiquidityTier); ok {
		r0 = rf(ctx, id, name, initialMarginPpm, maintenanceFractionPpm, impactNotional)
	} else {
		r0 = ret.Get(0).(perpetualstypes.LiquidityTier)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint32, string, uint32, uint32, uint64) error); ok {
		r1 = rf(ctx, id, name, initialMarginPpm, maintenanceFractionPpm, impactNotional)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetParams provides a mock function with given fields: ctx, params
func (_m *PerpetualsKeeper) SetParams(ctx types.Context, params perpetualstypes.Params) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, perpetualstypes.Params) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPerpetualsKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewPerpetualsKeeper creates a new instance of PerpetualsKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPerpetualsKeeper(t mockConstructorTestingTNewPerpetualsKeeper) *PerpetualsKeeper {
	mock := &PerpetualsKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

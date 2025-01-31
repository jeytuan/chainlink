// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/v2/common/types"
)

// FeeEstimator is an autogenerated mock type for the FeeEstimator type
type FeeEstimator[H types.Head[BLOCK_HASH], F txmgrtypes.Fee, MAXPRICE interface{}, BLOCK_HASH types.Hashable] struct {
	mock.Mock
}

// BumpFee provides a mock function with given fields: ctx, originalFee, feeLimit, maxFeePrice, attempts
func (_m *FeeEstimator[H, F, MAXPRICE, BLOCK_HASH]) BumpFee(ctx context.Context, originalFee F, feeLimit uint32, maxFeePrice MAXPRICE, attempts []txmgrtypes.PriorAttempt[F, BLOCK_HASH]) (F, uint32, error) {
	ret := _m.Called(ctx, originalFee, feeLimit, maxFeePrice, attempts)

	var r0 F
	var r1 uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, F, uint32, MAXPRICE, []txmgrtypes.PriorAttempt[F, BLOCK_HASH]) (F, uint32, error)); ok {
		return rf(ctx, originalFee, feeLimit, maxFeePrice, attempts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, F, uint32, MAXPRICE, []txmgrtypes.PriorAttempt[F, BLOCK_HASH]) F); ok {
		r0 = rf(ctx, originalFee, feeLimit, maxFeePrice, attempts)
	} else {
		r0 = ret.Get(0).(F)
	}

	if rf, ok := ret.Get(1).(func(context.Context, F, uint32, MAXPRICE, []txmgrtypes.PriorAttempt[F, BLOCK_HASH]) uint32); ok {
		r1 = rf(ctx, originalFee, feeLimit, maxFeePrice, attempts)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	if rf, ok := ret.Get(2).(func(context.Context, F, uint32, MAXPRICE, []txmgrtypes.PriorAttempt[F, BLOCK_HASH]) error); ok {
		r2 = rf(ctx, originalFee, feeLimit, maxFeePrice, attempts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Close provides a mock function with given fields:
func (_m *FeeEstimator[H, F, MAXPRICE, BLOCK_HASH]) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFee provides a mock function with given fields: ctx, calldata, feeLimit, maxFeePrice, opts
func (_m *FeeEstimator[H, F, MAXPRICE, BLOCK_HASH]) GetFee(ctx context.Context, calldata []byte, feeLimit uint32, maxFeePrice MAXPRICE, opts ...txmgrtypes.Opt) (F, uint32, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, calldata, feeLimit, maxFeePrice)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 F
	var r1 uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, uint32, MAXPRICE, ...txmgrtypes.Opt) (F, uint32, error)); ok {
		return rf(ctx, calldata, feeLimit, maxFeePrice, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, uint32, MAXPRICE, ...txmgrtypes.Opt) F); ok {
		r0 = rf(ctx, calldata, feeLimit, maxFeePrice, opts...)
	} else {
		r0 = ret.Get(0).(F)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, uint32, MAXPRICE, ...txmgrtypes.Opt) uint32); ok {
		r1 = rf(ctx, calldata, feeLimit, maxFeePrice, opts...)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	if rf, ok := ret.Get(2).(func(context.Context, []byte, uint32, MAXPRICE, ...txmgrtypes.Opt) error); ok {
		r2 = rf(ctx, calldata, feeLimit, maxFeePrice, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// HealthReport provides a mock function with given fields:
func (_m *FeeEstimator[H, F, MAXPRICE, BLOCK_HASH]) HealthReport() map[string]error {
	ret := _m.Called()

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *FeeEstimator[H, F, MAXPRICE, BLOCK_HASH]) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *FeeEstimator[H, F, MAXPRICE, BLOCK_HASH]) OnNewLongestChain(ctx context.Context, head H) {
	_m.Called(ctx, head)
}

// Ready provides a mock function with given fields:
func (_m *FeeEstimator[H, F, MAXPRICE, BLOCK_HASH]) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *FeeEstimator[H, F, MAXPRICE, BLOCK_HASH]) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFeeEstimator interface {
	mock.TestingT
	Cleanup(func())
}

// NewFeeEstimator creates a new instance of FeeEstimator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFeeEstimator[H types.Head[BLOCK_HASH], F txmgrtypes.Fee, MAXPRICE interface{}, BLOCK_HASH types.Hashable](t mockConstructorTestingTNewFeeEstimator) *FeeEstimator[H, F, MAXPRICE, BLOCK_HASH] {
	mock := &FeeEstimator[H, F, MAXPRICE, BLOCK_HASH]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

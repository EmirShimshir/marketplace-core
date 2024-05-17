// Code generated by mockery v2.42.2. DO NOT EDIT.

package repoMocks

import (
	context "context"

	"github.com/EmirShimshir/marketplace-core/domain"
	mock "github.com/stretchr/testify/mock"
)

// WithdrawRepository is an autogenerated mock type for the IWithdrawRepository type
type WithdrawRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, withdraw
func (_m *WithdrawRepository) Create(ctx context.Context, withdraw domain.Withdraw) (domain.Withdraw, error) {
	ret := _m.Called(ctx, withdraw)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 domain.Withdraw
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Withdraw) (domain.Withdraw, error)); ok {
		return rf(ctx, withdraw)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Withdraw) domain.Withdraw); ok {
		r0 = rf(ctx, withdraw)
	} else {
		r0 = ret.Get(0).(domain.Withdraw)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Withdraw) error); ok {
		r1 = rf(ctx, withdraw)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, withdrawID
func (_m *WithdrawRepository) Delete(ctx context.Context, withdrawID domain.ID) error {
	ret := _m.Called(ctx, withdrawID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) error); ok {
		r0 = rf(ctx, withdrawID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, limit, offset
func (_m *WithdrawRepository) Get(ctx context.Context, limit int64, offset int64) ([]domain.Withdraw, error) {
	ret := _m.Called(ctx, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []domain.Withdraw
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) ([]domain.Withdraw, error)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) []domain.Withdraw); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Withdraw)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, WithdrawID
func (_m *WithdrawRepository) GetByID(ctx context.Context, WithdrawID domain.ID) (domain.Withdraw, error) {
	ret := _m.Called(ctx, WithdrawID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 domain.Withdraw
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) (domain.Withdraw, error)); ok {
		return rf(ctx, WithdrawID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) domain.Withdraw); ok {
		r0 = rf(ctx, WithdrawID)
	} else {
		r0 = ret.Get(0).(domain.Withdraw)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, WithdrawID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByShopID provides a mock function with given fields: ctx, shopID
func (_m *WithdrawRepository) GetByShopID(ctx context.Context, shopID domain.ID) ([]domain.Withdraw, error) {
	ret := _m.Called(ctx, shopID)

	if len(ret) == 0 {
		panic("no return value specified for GetByShopID")
	}

	var r0 []domain.Withdraw
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) ([]domain.Withdraw, error)); ok {
		return rf(ctx, shopID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) []domain.Withdraw); ok {
		r0 = rf(ctx, shopID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Withdraw)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, shopID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, withdraw
func (_m *WithdrawRepository) Update(ctx context.Context, withdraw domain.Withdraw) (domain.Withdraw, error) {
	ret := _m.Called(ctx, withdraw)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 domain.Withdraw
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Withdraw) (domain.Withdraw, error)); ok {
		return rf(ctx, withdraw)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Withdraw) domain.Withdraw); ok {
		r0 = rf(ctx, withdraw)
	} else {
		r0 = ret.Get(0).(domain.Withdraw)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Withdraw) error); ok {
		r1 = rf(ctx, withdraw)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewWithdrawRepository creates a new instance of WithdrawRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWithdrawRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *WithdrawRepository {
	mock := &WithdrawRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

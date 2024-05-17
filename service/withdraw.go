package service

import (
	"context"
	"github.com/EmirShimshir/marketplace-core/domain"
	"github.com/EmirShimshir/marketplace-core/port"
)

type WithdrawService struct {
	repo port.IWithdrawRepository
}

func NewWithdrawService(repo port.IWithdrawRepository) *WithdrawService {
	return &WithdrawService{
		repo: repo,
	}
}

func (w *WithdrawService) Get(ctx context.Context, limit, offset int64) ([]domain.Withdraw, error) {
	return w.repo.Get(ctx, limit, offset)
}

func (w *WithdrawService) GetByID(ctx context.Context, withdrawID domain.ID) (domain.Withdraw, error) {
	return w.repo.GetByID(ctx, withdrawID)
}

func (w *WithdrawService) GetByShopID(ctx context.Context, shopID domain.ID) ([]domain.Withdraw, error) {
	return w.repo.GetByShopID(ctx, shopID)
}

func (w *WithdrawService) Create(ctx context.Context, param port.CreateWithdrawParam) (domain.Withdraw, error) {
	if param.Sum < 1 {
		return domain.Withdraw{}, domain.ErrPrice
	}

	return w.repo.Create(ctx, domain.Withdraw{
		ID:      domain.NewID(),
		ShopID:  param.ShopID,
		Comment: param.Comment,
		Sum:     param.Sum,
		Status:  domain.WithdrawStatusStart,
	})
}

func (w *WithdrawService) Update(ctx context.Context, withdrawID domain.ID, param port.UpdateWithdrawParam) (domain.Withdraw, error) {
	wr, err := w.GetByID(ctx, withdrawID)
	if err != nil {
		return domain.Withdraw{}, err
	}

	if param.Status != nil {
		wr.Status = *param.Status
	}
	if param.Sum.Valid {
		wr.Sum = param.Sum.Int64
	}
	if param.Comment.Valid {
		wr.Comment = param.Comment.String
	}

	if wr.Sum < 1 {
		return domain.Withdraw{}, domain.ErrPrice
	}

	return w.repo.Update(ctx, wr)
}

func (w *WithdrawService) Delete(ctx context.Context, withdrawID domain.ID) error {
	return w.repo.Delete(ctx, withdrawID)
}

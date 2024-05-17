package service

import (
	"context"
	"github.com/EmirShimshir/marketplace-core/domain"
	"github.com/EmirShimshir/marketplace-core/port"
)

type UserService struct {
	userRepo port.IUserRepository
}

func NewUserService(userRepo port.IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) Get(ctx context.Context, limit, offset int64) ([]domain.User, error) {
	return u.userRepo.Get(ctx, limit, offset)
}

func (u *UserService) GetByID(ctx context.Context, userID domain.ID) (domain.User, error) {
	return u.userRepo.GetByID(ctx, userID)
}

func (u *UserService) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	return u.userRepo.GetByEmail(ctx, email)
}

func (u *UserService) Create(ctx context.Context, param port.CreateUserParam) (domain.User, error) {
	if param.Name == "" {
		return domain.User{}, domain.ErrName
	}
	if param.Surname == "" {
		return domain.User{}, domain.ErrSurname
	}

	return u.userRepo.Create(ctx, domain.User{
		ID:       domain.NewID(),
		CartID:   domain.NewID(),
		Name:     param.Name,
		Surname:  param.Surname,
		Phone:    param.Phone,
		Email:    param.Email,
		Password: param.Password,
		Role:     param.Role,
	})
}

func (u *UserService) Update(ctx context.Context, userID domain.ID,
	param port.UpdateUserParam) (domain.User, error) {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	if param.Name.Valid {
		user.Name = param.Name.String
	}
	if param.Surname.Valid {
		user.Surname = param.Surname.String
	}
	if param.Phone.Valid {
		user.Phone = param.Phone
	}

	if user.Name == "" {
		return domain.User{}, domain.ErrName
	}
	if user.Surname == "" {
		return domain.User{}, domain.ErrSurname
	}

	return u.userRepo.Update(ctx, user)
}

func (u *UserService) Delete(ctx context.Context, userID domain.ID) error {
	return u.userRepo.Delete(ctx, userID)
}

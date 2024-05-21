package service

import (
	"context"
	"github.com/EmirShimshir/marketplace-core/domain"
	"github.com/EmirShimshir/marketplace-core/port"
	log "github.com/sirupsen/logrus"
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
	users, err := u.userRepo.Get(ctx, limit, offset)
	if err != nil{
		log.WithFields(log.Fields{
			"from": "UserServiceGet",
		}).Error(err.Error())
		return nil, err
	}

	return users, nil
}

func (u *UserService) GetByID(ctx context.Context, userID domain.ID) (domain.User, error) {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil{
		log.WithFields(log.Fields{
			"from": "UserServiceGetByID",
		}).Error(err.Error())
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserService) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil{
		log.WithFields(log.Fields{
			"from": "UserServiceGetByEmail",
		}).Error(err.Error())
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserService) Create(ctx context.Context, param port.CreateUserParam) (domain.User, error) {
	if param.Name == "" {
		log.WithFields(log.Fields{
			"from": "UserServiceCreate",
		}).Error(domain.ErrName.Error())
		return domain.User{}, domain.ErrName
	}
	if param.Surname == "" {
		log.WithFields(log.Fields{
			"from": "UserServiceCreate",
		}).Error(domain.ErrSurname.Error())
		return domain.User{}, domain.ErrSurname
	}

	user, err := u.userRepo.Create(ctx, domain.User{
		ID:       domain.NewID(),
		CartID:   domain.NewID(),
		Name:     param.Name,
		Surname:  param.Surname,
		Phone:    param.Phone,
		Email:    param.Email,
		Password: param.Password,
		Role:     param.Role,
	})
	if err != nil{
		log.WithFields(log.Fields{
			"from": "UserServiceCreate",
		}).Error(err.Error())
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserService) Update(ctx context.Context, userID domain.ID,
	param port.UpdateUserParam) (domain.User, error) {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UserServiceUpdate",
		}).Error(err.Error())
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
		log.WithFields(log.Fields{
			"from": "UserServiceUpdate",
		}).Error(domain.ErrName.Error())
		return domain.User{}, domain.ErrName
	}
	if user.Surname == "" {
		log.WithFields(log.Fields{
			"from": "UserServiceUpdate",
		}).Error(domain.ErrSurname.Error())
		return domain.User{}, domain.ErrSurname
	}

	user, err = u.userRepo.Update(ctx, user)
	if err != nil{
		log.WithFields(log.Fields{
			"from": "UserServiceUpdate",
		}).Error(err.Error())
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserService) Delete(ctx context.Context, userID domain.ID) error {
	err := u.userRepo.Delete(ctx, userID)
	if err != nil{
		log.WithFields(log.Fields{
			"from": "UserServiceDelete",
		}).Error(err.Error())
		return err
	}

	return nil
}

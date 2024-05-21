package service

import (
	"context"
	"github.com/EmirShimshir/marketplace-core/domain"
	"github.com/EmirShimshir/marketplace-core/port"
	log "github.com/sirupsen/logrus"
)

type AuthService struct {
	authProvider port.IAuthProvider
	userService  port.IUserService
}

func NewAuthService(authProvider port.IAuthProvider, userService port.IUserService) *AuthService {
	return &AuthService{
		authProvider: authProvider,
		userService:  userService,
	}
}

func (a *AuthService) SignIn(ctx context.Context, param port.SignInParam) (domain.AuthDetails, error) {
	user, err := a.userService.GetByEmail(ctx, param.Email)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "SignIn",
		}).Error(err.Error())
		return domain.AuthDetails{}, domain.ErrEmail
	}

	if user.Password != a.authProvider.GenPasswordHash(param.Password) {
		log.WithFields(log.Fields{
			"from": "SignIn",
		}).Error(domain.ErrPassword.Error())
		return domain.AuthDetails{}, domain.ErrPassword
	}
	ad, err :=  a.authProvider.CreateJWTSession(domain.AuthPayload{UserID: user.ID}, param.Fingerprint)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "SignIn",
		}).Error(err.Error())
		return domain.AuthDetails{}, err
	}

	return ad, nil
}

func (a *AuthService) SignUp(ctx context.Context, param port.SignUpParam) error {
	_, err := a.userService.Create(ctx, port.CreateUserParam{
		Name:     param.Name,
		Surname:  param.Surname,
		Phone:    param.Phone,
		Email:    param.Email,
		Password: a.authProvider.GenPasswordHash(param.Password),
		Role:     param.Role,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"from": "SignUp",
		}).Error(err.Error())
		return err
	}
	return nil
}

func (a *AuthService) LogOut(ctx context.Context, refreshToken domain.Token) error {
	err :=  a.authProvider.DeleteJWTSession(refreshToken)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "LogOut",
		}).Error(err.Error())
		return err
	}
	return nil
}

func (a *AuthService) Refresh(ctx context.Context, refreshToken domain.Token,
	fingerprint string) (domain.AuthDetails, error) {
	ad, err := a.authProvider.RefreshJWTSession(refreshToken, fingerprint)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "Refresh",
		}).Error(err.Error())
		return domain.AuthDetails{}, err
	}
	return ad, nil
}

func (a *AuthService) Verify(ctx context.Context, accessToken domain.Token) error {
	_, err := a.authProvider.VerifyJWTToken(accessToken)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "Verify",
		}).Error(err.Error())
		return err
	}
	return nil
}

func (a *AuthService) Payload(ctx context.Context, accessToken domain.Token) (domain.AuthPayload, error) {
	ap, err :=  a.authProvider.VerifyJWTToken(accessToken)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "Payload",
		}).Error(err.Error())
		return domain.AuthPayload{}, err
	}
	return ap, nil
}

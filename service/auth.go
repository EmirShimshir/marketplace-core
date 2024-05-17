package service

import (
	"context"
	"github.com/EmirShimshir/marketplace-core/domain"
	"github.com/EmirShimshir/marketplace-core/port"
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
		return domain.AuthDetails{}, domain.ErrEmail
	}

	if user.Password != a.authProvider.GenPasswordHash(param.Password) {
		return domain.AuthDetails{}, domain.ErrPassword
	}
	return a.authProvider.CreateJWTSession(domain.AuthPayload{UserID: user.ID}, param.Fingerprint)
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

	return err
}

func (a *AuthService) LogOut(ctx context.Context, refreshToken domain.Token) error {
	return a.authProvider.DeleteJWTSession(refreshToken)
}

func (a *AuthService) Refresh(ctx context.Context, refreshToken domain.Token,
	fingerprint string) (domain.AuthDetails, error) {
	return a.authProvider.RefreshJWTSession(refreshToken, fingerprint)
}

func (a *AuthService) Verify(ctx context.Context, accessToken domain.Token) error {
	_, err := a.authProvider.VerifyJWTToken(accessToken)
	return err
}

func (a *AuthService) Payload(ctx context.Context, accessToken domain.Token) (domain.AuthPayload, error) {
	return a.authProvider.VerifyJWTToken(accessToken)
}

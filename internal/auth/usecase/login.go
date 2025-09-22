package usecase

import (
	"context"
	"errors"

	"github.com/amagkn/sso-service/internal/auth/dto"
	"github.com/amagkn/sso-service/internal/auth/entity"
	"github.com/amagkn/sso-service/pkg/base_errors"
)

func (uc *UseCase) Login(ctx context.Context, input dto.LoginInput) (dto.LoginOutput, error) {
	var output dto.LoginOutput

	user, err := uc.postgres.SelectUser(ctx, input.Email)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			return output, base_errors.InvalidCredentials
		}

		return output, base_errors.WithPath("uc.postgres.SelectUser", err)
	}

	err = user.ComparePassword([]byte(input.Password))
	if err != nil {
		return output, base_errors.InvalidCredentials
	}

	app, err := uc.postgres.SelectApp(ctx, input.AppID)
	if err != nil {
		if errors.Is(err, entity.ErrAppNotFound) {
			return output, entity.ErrInvalidAppID
		}

		return output, base_errors.WithPath("uc.postgres.SelectApp", err)
	}

	output.Token, err = user.NewJWTToken(app, uc.tokenTTL)
	if err != nil {
		return output, base_errors.WithPath("user.NewJWTToken", err)
	}

	return output, nil
}

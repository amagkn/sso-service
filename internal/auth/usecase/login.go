package usecase

import (
	"context"

	"github.com/amagkn/sso-service/internal/auth/dto"
	"github.com/amagkn/sso-service/pkg/logger"
)

func (uc *UseCase) Login(ctx context.Context, input dto.LoginInput) (dto.LoginOutput, error) {
	var output dto.LoginOutput

	user, err := uc.postgres.SelectUserByEmail(ctx, input.Email)
	if err != nil {
		logger.Error(err, "uc.postgres.SelectUserByEmail")

		return output, err
	}

	err = user.ComparePassword([]byte(input.Password))
	if err != nil {
		logger.Error(err, "user.ComparePassword")

		return output, err
	}

	app, err := uc.postgres.SelectAppByID(ctx, input.AppID)
	if err != nil {
		logger.Error(err, "uc.postgres.SelectAppByID")

		return output, err
	}

	output.Token, err = user.NewJWTToken(app, uc.tokenTTL)
	if err != nil {
		logger.Error(err, "user.NewJWTToken")

		return output, err
	}

	return output, nil
}

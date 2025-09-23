package usecase

import (
	"context"

	"github.com/amagkn/sso-service/internal/auth/dto"
	"github.com/amagkn/sso-service/pkg/logger"
)

func (uc *UseCase) IsAdmin(ctx context.Context, input dto.IsAdminInput) (dto.IsAdminOutput, error) {
	var output dto.IsAdminOutput
	var err error

	output.IsAdmin, err = uc.postgres.UserIsAdmin(ctx, input.UserId)
	if err != nil {
		logger.Error(err, "uc.postgres.UserIsAdmin")

		return output, err
	}

	return output, nil
}

package usecase

import (
	"context"

	"github.com/amagkn/sso-service/internal/auth/dto"
	"github.com/amagkn/sso-service/pkg/base_errors"
)

func (uc *UseCase) IsAdmin(ctx context.Context, input dto.IsAdminInput) (dto.IsAdminOutput, error) {
	var output dto.IsAdminOutput
	var err error

	output.IsAdmin, err = uc.postgres.UserIsAdmin(ctx, input.UserId)
	if err != nil {
		return output, base_errors.WithPath("uc.postgres.UserIsAdmin", err)
	}

	return output, nil
}

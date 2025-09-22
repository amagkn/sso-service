package usecase

import (
	"context"

	"github.com/amagkn/sso-service/internal/auth/dto"
)

func (uc *UseCase) IsAdmin(ctx context.Context, input dto.IsAdminInput) (dto.IsAdminOutput, error) {
	return dto.IsAdminOutput{}, nil
}

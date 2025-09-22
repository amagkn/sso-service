package usecase

import (
	"context"

	"github.com/amagkn/sso-service/internal/auth/dto"
)

func (uc *UseCase) Register(ctx context.Context, input dto.RegisterInput) (dto.RegisterOutput, error) {
	return dto.RegisterOutput{}, nil
}

package usecase

import (
	"context"

	"github.com/amagkn/sso-service/internal/auth/dto"
)

func (uc *UseCase) Login(ctx context.Context, input dto.LoginInput) (dto.LoginOutput, error) {
	return dto.LoginOutput{}, nil
}

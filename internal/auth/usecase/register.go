package usecase

import (
	"context"

	"github.com/amagkn/sso-service/internal/auth/dto"
	"github.com/amagkn/sso-service/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) Register(ctx context.Context, input dto.RegisterInput) (dto.RegisterOutput, error) {
	var output dto.RegisterOutput

	passHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(err, "bcrypt.GenerateFromPassword")

		return output, err
	}

	saveUserDto := dto.InsertUserInput{Email: input.Email, PassHash: passHash}
	userId, err := uc.postgres.InsertUser(ctx, saveUserDto)
	if err != nil {
		logger.Error(err, "uc.postgres.InsertUser")

		return output, err
	}

	output.UserId = userId

	return output, nil
}

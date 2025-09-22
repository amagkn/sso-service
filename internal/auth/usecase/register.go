package usecase

import (
	"context"
	"errors"

	"github.com/amagkn/sso-service/internal/auth/dto"
	"github.com/amagkn/sso-service/internal/auth/entity"
	"github.com/amagkn/sso-service/pkg/base_errors"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) Register(ctx context.Context, input dto.RegisterInput) (dto.RegisterOutput, error) {
	var output dto.RegisterOutput

	passHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return output, base_errors.WithPath("bcrypt.GenerateFromPassword", err)
	}

	saveUserDto := dto.InsertUserInput{Email: input.Email, PassHash: passHash}

	user, err := uc.postgres.InsertUser(ctx, saveUserDto)
	if err != nil {
		if errors.Is(err, entity.ErrUserAlreadyExists) {
			return output, entity.ErrUserAlreadyExists
		}

		return output, base_errors.WithPath("uc.postgres.InsertUser", err)
	}

	output.UserId = user.ID

	return output, nil
}

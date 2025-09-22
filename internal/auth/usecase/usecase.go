package usecase

import (
	"context"
	"time"

	"github.com/amagkn/sso-service/internal/auth/dto"
	"github.com/amagkn/sso-service/internal/auth/entity"
)

type Postgres interface {
	InsertUser(ctx context.Context, input dto.InsertUserInput) (entity.User, error)
	SelectUser(ctx context.Context, email string) (entity.User, error)
	SelectApp(ctx context.Context, appID int32) (entity.App, error)
}

type UseCase struct {
	postgres Postgres
	tokenTTL time.Duration
}

func New(p Postgres, tokenTTL time.Duration) *UseCase {
	return &UseCase{postgres: p, tokenTTL: tokenTTL}
}

package postgres

import (
	"context"

	"github.com/amagkn/sso-service/internal/auth/dto"
	"github.com/amagkn/sso-service/internal/auth/entity"
)

func (p *Postgres) InsertUser(ctx context.Context, input dto.InsertUserInput) (entity.User, error) {
	return entity.User{}, nil
}

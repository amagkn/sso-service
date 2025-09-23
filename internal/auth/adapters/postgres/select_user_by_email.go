package postgres

import (
	"context"

	"github.com/amagkn/sso-service/internal/auth/entity"
)

func (p *Postgres) SelectUserByEmail(ctx context.Context, email string) (entity.User, error) {
	return entity.User{}, nil
}

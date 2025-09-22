package postgres

import (
	"context"

	"github.com/amagkn/sso-service/internal/auth/entity"
)

func (p *Postgres) SelectApp(ctx context.Context, appID int32) (entity.App, error) {
	return entity.App{}, nil
}

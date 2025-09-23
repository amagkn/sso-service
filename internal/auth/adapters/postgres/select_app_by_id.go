package postgres

import (
	"context"
	"errors"

	"github.com/amagkn/sso-service/internal/auth/entity"
	"github.com/amagkn/sso-service/pkg/base_errors"
	"github.com/amagkn/sso-service/pkg/postgres"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) SelectAppByID(ctx context.Context, appID int32) (entity.App, error) {
	var app entity.App

	// SELECT id, name, secret FROM app WHERE id = {appID}
	ds := goqu.From("app").
		Select("id", "name", "secret").
		Where(goqu.Ex{"id": appID})

	sql, _, err := ds.ToSQL()
	if err != nil {
		return app, base_errors.WithPath("ds.ToSQL", err)
	}

	err = p.pool.QueryRow(ctx, sql).Scan(&app.ID, &app.Name, &app.Secret)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return app, entity.ErrAppNotFound
		}

		return app, base_errors.WithPath("p.pool.QueryRow.Scan", err)
	}

	return app, nil
}

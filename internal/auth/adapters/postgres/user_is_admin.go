package postgres

import (
	"context"
	"errors"

	"github.com/amagkn/sso-service/internal/auth/entity"
	"github.com/amagkn/sso-service/pkg/base_errors"
	"github.com/amagkn/sso-service/pkg/postgres"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) UserIsAdmin(ctx context.Context, id int64) (bool, error) {
	var isAdmin bool

	// SELECT is_admin FROM "user" WHERE id = {id} LIMIT 1
	ds := goqu.From("user").Select("is_admin").Where(goqu.Ex{"id": id}).Limit(1)

	sql, _, err := ds.ToSQL()
	if err != nil {
		return isAdmin, base_errors.WithPath("ds.ToSQL", err)
	}

	err = p.pool.QueryRow(ctx, sql).Scan(&isAdmin)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return isAdmin, entity.ErrUserNotFound
		}

		return isAdmin, base_errors.WithPath("p.pool.QueryRow.Scan", err)
	}

	return isAdmin, nil
}

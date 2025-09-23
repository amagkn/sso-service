package postgres

import (
	"context"

	"github.com/amagkn/sso-service/internal/auth/dto"
	"github.com/amagkn/sso-service/internal/auth/entity"
	"github.com/amagkn/sso-service/pkg/base_errors"
	"github.com/amagkn/sso-service/pkg/postgres"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) InsertUser(ctx context.Context, input dto.InsertUserInput) (int64, error) {
	var userId int64

	ds := goqu.Insert("user").
		Rows(goqu.Record{
			"email":     input.Email,
			"pass_hash": input.PassHash,
		}).
		Returning("id")

	sql, args, err := ds.ToSQL()
	if err != nil {
		return userId, base_errors.WithPath("ds.ToSQL", err)
	}

	row := p.pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&userId)
	if err != nil {
		if postgres.IsUniqueConstraintError(err) {
			return userId, entity.ErrUserAlreadyExists
		}

		return userId, base_errors.WithPath("row.Scan", err)
	}

	return userId, nil
}

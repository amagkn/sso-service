package postgres

import (
	"context"
	"errors"

	"github.com/amagkn/sso-service/internal/auth/entity"
	"github.com/amagkn/sso-service/pkg/base_errors"
	"github.com/amagkn/sso-service/pkg/postgres"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) SelectUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User

	// SELECT id, email, pass_hash, is_admin from "user" WHERE email = {email}
	ds := goqu.From("user").
		Select("id", "email", "pass_hash", "is_admin").
		Where(goqu.Ex{"email": email})

	sql, _, err := ds.ToSQL()
	if err != nil {
		return user, base_errors.WithPath("ds.ToSQL", err)
	}

	err = p.pool.QueryRow(ctx, sql).Scan(&user.ID, &user.Email, &user.PassHash, &user.IsAdmin)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return user, entity.ErrUserNotFound
		}

		return user, base_errors.WithPath("p.pool.QueryRow.Scan", err)
	}

	return user, nil
}

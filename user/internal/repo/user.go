package repo

import (
	"context"
	"database/sql"
	"user/internal/model"
)

type UserRepo interface {
	Save(tx *sql.Tx, ctx context.Context, user *model.User) error
	FindById(ctx context.Context, id int64) (user *model.User, err error)
}

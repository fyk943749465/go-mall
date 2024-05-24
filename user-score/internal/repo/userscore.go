package repo

import (
	"context"
	"database/sql"
	"user-score/internal/model"
)

type UserScoreRepo interface {
	SaveUserScore(tx *sql.Tx, ctx context.Context, score *model.UserScore) error
}

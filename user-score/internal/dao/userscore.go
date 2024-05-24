package dao

import (
	"context"
	"database/sql"
	"fmt"
	"user-score/database"
	"user-score/internal/model"
)

var cacheUserIdPrefix = "cache:userscore:id:"

type UserScoreDao struct {
	*database.DBConn
}

func NewUserScoreDao(conn *database.DBConn) *UserScoreDao {
	return &UserScoreDao{
		conn,
	}
}

func (d *UserScoreDao) SaveUserScore(tx *sql.Tx, ctx context.Context, score *model.UserScore) error {
	sql := fmt.Sprintf("insert into %s (user_id, score) values(?, ?)", score.TableName())
	result, err := tx.ExecContext(ctx, sql, score.UserId, score.Score)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	score.Id = id
	return nil

}

package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/database"
	"user/internal/model"
)

var cacheUserIdPrefix = "cache:user:id:"

type UserDao struct {
	*database.DBConn
}

func (d *UserDao) FindById(ctx context.Context, id int64) (user *model.User, err error) {

	user = &model.User{}
	querySql := fmt.Sprintf("select * from %s where id = ?", user.TableName())
	userIdKey := fmt.Sprintf("%s%d", cacheUserIdPrefix, id)
	err = d.ConnCache.QueryRowCtx(ctx, user, userIdKey,
		func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return conn.QueryRowCtx(ctx, v, querySql, id)
		})
	return
}

func NewUserDao(conn *database.DBConn) *UserDao {
	return &UserDao{
		conn,
	}
}

func (d *UserDao) Save(tx *sql.Tx, ctx context.Context, user *model.User) error {
	sql := fmt.Sprintf("insert into %s (id, name, gender) values(?, ?, ?)", user.TableName())
	result, err := tx.ExecContext(ctx, sql, user.Id, user.Name, user.Gender)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = id
	return nil

}

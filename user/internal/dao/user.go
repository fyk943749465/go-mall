package dao

import (
	"context"
	"fmt"
	"user/database"
	"user/internal/model"
)

type UserDao struct {
	*database.DBConn
}

func NewUserDao(conn *database.DBConn) *UserDao {
	return &UserDao{
		conn,
	}
}

func (d *UserDao) Save(ctx context.Context, user *model.User) error {
	sql := fmt.Sprintf("insert into %s (name, gender) values(?, ?)", user.TableName())
	result, err := d.Conn.ExecCtx(ctx, sql, user.Name, user.Gender)
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

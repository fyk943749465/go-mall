package database

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// we use go-zero sqlx

type DBConn struct {
	Conn      sqlx.SqlConn    // mysql
	ConnCache sqlc.CachedConn // redis
}

func Connect(datasource string, conf cache.CacheConf) *DBConn {
	sqlConn := sqlx.NewMysql(datasource)
	d := &DBConn{
		Conn: sqlConn,
	}
	if conf != nil {
		cachedConn := sqlc.NewConn(sqlConn, conf)
		d.ConnCache = cachedConn
	}
	return d
}

package svc

import (
	"database/sql"
	"user-score/database"
	"user-score/internal/config"
	"user-score/internal/dao"
	"user-score/internal/repo"
)

type ServiceContext struct {
	Config        config.Config
	UserScoreRepo repo.UserScoreRepo
	Db            *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	connect := database.Connect(c.Mysql.DataSource, c.CacheRedis)
	db, _ := connect.Conn.RawDB()
	return &ServiceContext{
		Config:        c,
		UserScoreRepo: dao.NewUserScoreDao(connect),
		Db:            db,
	}
}

package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"micro-ping/service/model"
	"micro-ping/service/user/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	Model       model.TbUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(c.Redis.RedisConf),
		Model:       model.NewTbUserModel(sqlx.NewMysql(c.Datasource), c.Cache),
	}
}

package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"micro-ping/service/model"
	"micro-ping/service/shop/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Model  model.TbShopModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewTbShopModel(sqlx.NewMysql(c.Datasource), c.Cache),
	}
}

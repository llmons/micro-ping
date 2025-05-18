package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"micro-ping/service/model"
	"shop_type/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Model  model.TbShopTypeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewTbShopTypeModel(sqlx.NewMysql(c.Datasource), c.Cache),
	}
}

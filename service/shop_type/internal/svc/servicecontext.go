package svc

import (
	"micro-ping/service/model"
	"shop_type/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
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

package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"restful/internal/config"
	"shop_type/shoptypeclient"
)

type ServiceContext struct {
	Config   config.Config
	ShopType shoptypeclient.ShopTypeZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		ShopType: shoptypeclient.NewShopTypeZrpcClient(zrpc.MustNewClient(c.ShopType)),
	}
}

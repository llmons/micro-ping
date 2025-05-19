package svc

import (
	"restful/internal/config"
	"shop_type/shoptypeclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	ShopTypeRpc shoptypeclient.ShopTypeZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		ShopTypeRpc: shoptypeclient.NewShopTypeZrpcClient(zrpc.MustNewClient(c.ShopTypeRpcConf)),
	}
}

package svc

import (
	"micro-ping/restful/internal/config"
	"micro-ping/service/shop_type/shoptypeclient"
	"micro-ping/service/user/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	ShopTypeRpc shoptypeclient.ShopTypeZrpcClient
	UserRpc     userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		ShopTypeRpc: shoptypeclient.NewShopTypeZrpcClient(zrpc.MustNewClient(c.ShopTypeRpcConf)),
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}

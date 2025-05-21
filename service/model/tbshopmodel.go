package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbShopModel = (*customTbShopModel)(nil)

type (
	// TbShopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbShopModel.
	TbShopModel interface {
		tbShopModel
	}

	customTbShopModel struct {
		*defaultTbShopModel
	}
)

// NewTbShopModel returns a model for the database table.
func NewTbShopModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbShopModel {
	return &customTbShopModel{
		defaultTbShopModel: newTbShopModel(conn, c, opts...),
	}
}

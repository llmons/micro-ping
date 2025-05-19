package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbShopTypeModel = (*customTbShopTypeModel)(nil)

type (
	// TbShopTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbShopTypeModel.
	TbShopTypeModel interface {
		tbShopTypeModel
		RetrieveAll() ([]*TbShopType, error)
	}

	customTbShopTypeModel struct {
		*defaultTbShopTypeModel
	}
)

// NewTbShopTypeModel returns a model for the database table.
func NewTbShopTypeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbShopTypeModel {
	return &customTbShopTypeModel{
		defaultTbShopTypeModel: newTbShopTypeModel(conn, c, opts...),
	}
}

func (m *defaultTbShopTypeModel) RetrieveAll() ([]*TbShopType, error) {
	var resp []*TbShopType
	err := m.QueryRowsNoCache(&resp, "select "+tbShopTypeRows+" from "+m.table)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

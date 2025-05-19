package logic

import (
	"context"

	"micro-ping/service/shop_type/internal/svc"
	"micro-ping/service/shop_type/shop_type"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTypeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTypeListLogic {
	return &GetTypeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTypeListLogic) GetTypeList(in *shop_type.ReqShopType) (*shop_type.RespShopType, error) {
	typeList, err := l.svcCtx.Model.RetrieveAll()
	if err != nil {
		l.Logger.Errorf("failed to retrieve type list: %v", err)
		return nil, err
	}

	resp := &shop_type.RespShopType{}
	for _, t := range typeList {
		resp.ShopTypeList = append(resp.ShopTypeList, &shop_type.ShopType{
			Id:   t.Id,
			Name: t.Name.String,
			Icon: t.Icon.String,
			Sort: t.Sort.Int64,
		})
	}
	return resp, nil
}

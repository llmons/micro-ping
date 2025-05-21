package logic

import (
	"context"
	"micro-ping/service/shop/internal/svc"
	"micro-ping/service/shop/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopByIdLogic {
	return &GetShopByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetShopByIdLogic) GetShopById(in *shop.ReqGetShopById) (*shop.RespGetShopById, error) {
	one, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Errorf("get shop by id error: %v", err)
		return nil, err
	}

	return &shop.RespGetShopById{
		Shop: &shop.Shop{
			Name:      one.Name,
			TypeId:    one.TypeId,
			Images:    one.Images,
			Area:      one.Area.String,
			Address:   one.Address,
			X:         one.X,
			Y:         one.Y,
			AvgPrice:  one.AvgPrice.Int64,
			Sold:      one.Sold,
			Comments:  one.Comments,
			Score:     one.Score,
			OpenHours: one.OpenHours.String,
		},
	}, nil
}

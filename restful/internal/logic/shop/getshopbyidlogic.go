package shop

import (
	"context"
	"micro-ping/service/shop/shops"

	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopByIdLogic {
	return &GetShopByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShopByIdLogic) GetShopById(req *types.ReqGetShopById) (resp *types.RespGetShopById, err error) {
	rpcResp, err := l.svcCtx.ShopRpc.GetShopById(l.ctx, &shops.ReqGetShopById{Id: req.Id})
	if err != nil {
		l.Errorf("get shop by id error: %v", err)
		return
	}

	resp = &types.RespGetShopById{
		Shop: types.Shop{
			Name:      rpcResp.Shop.Name,
			TypeId:    rpcResp.Shop.TypeId,
			Images:    rpcResp.Shop.Images,
			Area:      rpcResp.Shop.Area,
			Address:   rpcResp.Shop.Address,
			X:         rpcResp.Shop.X,
			Y:         rpcResp.Shop.Y,
			AvgPrice:  rpcResp.Shop.AvgPrice,
			Sold:      rpcResp.Shop.Sold,
			Comments:  rpcResp.Shop.Comments,
			Score:     rpcResp.Shop.Score,
			OpenHours: rpcResp.Shop.OpenHours,
		},
	}
	return
}

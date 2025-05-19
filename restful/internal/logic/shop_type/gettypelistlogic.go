package shop_type

import (
	"context"
	"shop_type/shoptypeclient"

	"restful/internal/svc"
	"restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTypeListLogic {
	return &GetTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTypeListLogic) GetTypeList(req *types.ReqGetTypeList) (resp *types.RespGetTypeList, err error) {
	rpcResp, err := l.svcCtx.ShopTypeRpc.GetTypeList(l.ctx, &shoptypeclient.ReqShopType{})
	if err != nil {
		l.Logger.Errorf("get type list failed, err: %v", err)
		return nil, err
	}

	typeList := make([]*types.ShopType, 0)
	for _, shopType := range rpcResp.ShopTypeList {
		typeList = append(typeList, &types.ShopType{
			Id:   shopType.Id,
			Name: shopType.Name,
			Icon: shopType.Icon,
			Sort: shopType.Sort,
		})
	}
	resp = &types.RespGetTypeList{TypeList: typeList}
	return
}

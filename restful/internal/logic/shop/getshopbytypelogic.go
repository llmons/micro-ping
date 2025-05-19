package shop

import (
	"context"

	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopByTypeLogic {
	return &GetShopByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShopByTypeLogic) GetShopByType(req *types.ReqGetShopByType) (resp *types.RespGetShopByType, err error) {
	// todo: add your logic here and delete this line

	return
}

package shop

import (
	"context"

	"restful/internal/svc"
	"restful/internal/types"

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
	// todo: add your logic here and delete this line

	return
}

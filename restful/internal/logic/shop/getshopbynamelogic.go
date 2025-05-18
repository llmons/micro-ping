package shop

import (
	"context"

	"restful/internal/svc"
	"restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopByNameLogic {
	return &GetShopByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShopByNameLogic) GetShopByName(req *types.ReqGetShopByName) (resp *types.RespGetShopByName, err error) {
	// todo: add your logic here and delete this line

	return
}

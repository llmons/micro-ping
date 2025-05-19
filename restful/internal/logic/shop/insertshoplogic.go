package shop

import (
	"context"

	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertShopLogic {
	return &InsertShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertShopLogic) InsertShop(req *types.ReqInsertShop) (resp *types.RespInsertShop, err error) {
	// todo: add your logic here and delete this line

	return
}

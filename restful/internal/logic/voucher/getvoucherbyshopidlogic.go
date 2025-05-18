package voucher

import (
	"context"

	"restful/internal/svc"
	"restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVoucherByShopIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVoucherByShopIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVoucherByShopIdLogic {
	return &GetVoucherByShopIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVoucherByShopIdLogic) GetVoucherByShopId(req *types.ReqGetVoucherByShopId) (resp *types.RespGetVoucherByShopId, err error) {
	// todo: add your logic here and delete this line

	return
}

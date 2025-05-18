package voucher

import (
	"context"

	"restful/internal/svc"
	"restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSeckillVoucherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSeckillVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSeckillVoucherLogic {
	return &AddSeckillVoucherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSeckillVoucherLogic) AddSeckillVoucher(req *types.ReqAddSeckillVoucher) (resp *types.RespAddSeckillVoucher, err error) {
	// todo: add your logic here and delete this line

	return
}

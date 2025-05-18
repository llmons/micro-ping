package voucher

import (
	"context"

	"restful/internal/svc"
	"restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVoucherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVoucherLogic {
	return &AddVoucherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddVoucherLogic) AddVoucher(req *types.ReqAddVoucher) (resp *types.RespAddVoucher, err error) {
	// todo: add your logic here and delete this line

	return
}

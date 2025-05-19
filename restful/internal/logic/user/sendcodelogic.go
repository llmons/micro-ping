package user

import (
	"context"

	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"
	"micro-ping/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.ReqSendCode) (resp *types.RespSendCode, err error) {
	rpcResp, err := l.svcCtx.UserRpc.SendCode(l.ctx, &user.ReqSendCode{
		Phone: req.Phone,
	})
	if err != nil {
		l.Errorf("SendCode rpc error: %v", err)
		return nil, err
	}

	resp = &types.RespSendCode{
		Code: rpcResp.Code,
	}
	return
}

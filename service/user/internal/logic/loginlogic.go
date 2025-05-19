package logic

import (
	"context"

	"micro-ping/service/user/internal/svc"
	"micro-ping/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.ReqLogin) (*user.RespLogin, error) {
	// todo: add your logic here and delete this line

	return &user.RespLogin{}, nil
}

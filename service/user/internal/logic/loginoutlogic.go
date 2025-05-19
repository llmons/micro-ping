package logic

import (
	"context"

	"micro-ping/service/user/internal/svc"
	"micro-ping/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginOutLogic {
	return &LoginOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginOutLogic) LoginOut(in *user.ReqLogout) (*user.RespLogout, error) {
	// todo: add your logic here and delete this line

	return &user.RespLogout{}, nil
}

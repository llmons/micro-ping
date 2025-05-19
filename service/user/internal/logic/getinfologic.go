package logic

import (
	"context"

	"micro-ping/service/user/internal/svc"
	"micro-ping/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoLogic) GetInfo(in *user.ReqGetInfo) (*user.RespGetInfo, error) {
	// todo: add your logic here and delete this line

	return &user.RespGetInfo{}, nil
}

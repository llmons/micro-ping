package logic

import (
	"context"

	"micro-ping/service/user/internal/svc"
	"micro-ping/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMeLogic {
	return &GetMeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMeLogic) GetMe(in *user.ReqGetMe) (*user.RespGetMe, error) {
	one, err := l.svcCtx.Model.FindOne(l.ctx, in.UserId)
	if err != nil {
		l.Errorf("failed to find user: %v", err)
		return nil, err
	}

	return &user.RespGetMe{
		Phone:    one.Phone,
		Nickname: one.NickName,
		Icon:     one.Icon,
	}, nil
}

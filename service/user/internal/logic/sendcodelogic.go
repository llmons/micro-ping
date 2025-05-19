package logic

import (
	"context"
	"errors"
	"regexp"

	"micro-ping/service/user/internal/svc"
	"micro-ping/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendCodeLogic) SendCode(in *user.ReqSendCode) (*user.RespSendCode, error) {
	// check phone
	pattern := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !pattern.MatchString(in.Phone) {
		l.Errorf("invalid phone number: %s", in.Phone)
		return nil, errors.New("invalid phone number")
	}
	return &user.RespSendCode{}, nil
}

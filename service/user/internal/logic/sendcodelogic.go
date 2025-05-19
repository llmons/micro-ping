package logic

import (
	"context"
	"errors"
	"math/rand"
	"micro-ping/service/user/internal/base/constants"
	"regexp"
	"strconv"

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

	// check exist
	key := constants.RedisKeyPhonePrefix + in.Phone
	exist, err := l.svcCtx.RedisClient.Exists(key)
	if err != nil {
		l.Errorf("failed to check phone: %v", err)
		return nil, err
	}
	if exist {
		record, err := l.svcCtx.RedisClient.Get(key)
		if err != nil {
			l.Errorf("failed to get phone record: %v", err)
			return nil, err
		}
		return &user.RespSendCode{
			Code: record,
		}, nil
	}

	// new code
	code := strconv.Itoa(rand.Intn(99_999) + 100_000)
	err = l.svcCtx.RedisClient.Setex(key, code, constants.RedisTtlPhone)
	if err != nil {
		l.Errorf("failed to set phone code: %v", err)
		return nil, err
	}

	return &user.RespSendCode{
		Code: code,
	}, nil
}

package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stringx"
	"micro-ping/service/model"
	"micro-ping/service/user/internal/base/constants"
	"micro-ping/service/user/internal/svc"
	"micro-ping/service/user/user"
	"regexp"

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
	// check phone
	pattern := `^1[3-9]\d{9}$`
	if !regexp.MustCompile(pattern).MatchString(in.Phone) {
		l.Errorf("Invalid phone number: %s", in.Phone)
		return nil, errors.New("invalid phone number")
	}

	// check code match phone
	key := constants.RedisKeyPhonePrefix + in.Phone
	record, err := l.svcCtx.RedisClient.Get(key)
	if err != nil {
		l.Errorf("Redis get error: %v", err)
		return nil, err
	}

	fmt.Println(record)

	if record == "" || record != in.Code {
		l.Errorf("Code mismatch: %s != %s", record, in.Code)
		return nil, errors.New("code mismatch")
	}

	// matched
	// check user exists
	userRecord, err := l.svcCtx.Model.FindOneByPhone(l.ctx, in.Phone)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		l.Errorf("FindOneByPhone error: %v", err)
		return nil, err
	}

	if userRecord != nil {
		return &user.RespLogin{
			UserId: userRecord.Id,
		}, nil
	}

	// register user
	result, err := l.svcCtx.Model.Insert(l.ctx, &model.TbUser{
		Phone:    in.Phone,
		Password: stringx.Randn(16),
		NickName: stringx.Randn(5),
		Icon:     stringx.Randn(8),
	})
	if err != nil {
		l.Errorf("Insert error: %v", err)
		return nil, err
	}

	useId, err := result.LastInsertId()
	if err != nil {
		l.Errorf("LastInsertId error: %v", err)
		return nil, err
	}
	return &user.RespLogin{
		UserId: uint64(useId),
	}, nil
}

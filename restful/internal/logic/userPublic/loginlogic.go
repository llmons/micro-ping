package userPublic

import (
	"context"
	"micro-ping/restful/internal/base/jwt"
	"micro-ping/service/user/user"
	"time"

	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.ReqLogin) (resp *types.RespLogin, err error) {
	// check match
	rpcResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.ReqLogin{Phone: req.Phone, Code: req.Code})
	if err != nil {
		l.Errorf("Login rpc error: %v", err)
		return nil, err
	}

	var accessExpire = l.svcCtx.Config.JwtAuth.AccessExpire

	now := time.Now().Unix()
	preloads := make(map[string]interface{})
	preloads["user_id"] = rpcResp.UserId
	accessToken, err := jwt.GenToken(now, l.svcCtx.Config.JwtAuth.AccessSecret, nil, accessExpire)
	if err != nil {
		return nil, err
	}

	resp = &types.RespLogin{
		AccessToken:  accessToken,
		AccessExpire: accessExpire,
	}
	return
}

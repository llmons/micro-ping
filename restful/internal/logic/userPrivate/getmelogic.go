package userPrivate

import (
	"context"
	"micro-ping/restful/internal/base/jwt"
	"micro-ping/service/user/user"
	"strings"

	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMeLogic {
	return &GetMeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMeLogic) GetMe(req *types.ReqGetMe) (resp *types.RespGetMe, err error) {
	tokenStr, found := strings.CutPrefix(req.Autherization, "Bearer ")
	if !found {
		l.Errorf("Authorization header is not in Bearer format")
		return nil, err
	}
	claims, err := jwt.ParseToken(tokenStr, l.svcCtx.Config.JwtAuth.AccessSecret)
	if err != nil {
		return nil, err
	}

	userId := claims["user_id"].(float64)

	me, err := l.svcCtx.UserRpc.GetMe(l.ctx, &user.ReqGetMe{
		UserId: uint64(userId),
	})
	if err != nil {
		l.Errorf("GetMe rpc error: %v", err)
		return nil, err
	}
	resp = &types.RespGetMe{
		Phone:    me.Phone,
		Nickname: me.Nickname,
		Icon:     me.Icon,
	}
	return
}

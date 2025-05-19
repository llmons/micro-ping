package blog

import (
	"context"

	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlogLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBlogLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlogLikeLogic {
	return &BlogLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BlogLikeLogic) BlogLike(req *types.ReqLikeBlog) (resp *types.RespLikeBlog, err error) {
	// todo: add your logic here and delete this line

	return
}

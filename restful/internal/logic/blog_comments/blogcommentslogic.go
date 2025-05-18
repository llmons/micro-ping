package blog_comments

import (
	"context"

	"restful/internal/svc"
	"restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlogCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBlogCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlogCommentsLogic {
	return &BlogCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BlogCommentsLogic) BlogComments(req *types.ReqBlogComments) (resp *types.RespBlogComments, err error) {
	// todo: add your logic here and delete this line

	return
}

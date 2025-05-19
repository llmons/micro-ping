package blog

import (
	"context"

	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyBlogLogic {
	return &GetMyBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyBlogLogic) GetMyBlog(req *types.ReqGetMyBlog) (resp *types.RespGetMyBlog, err error) {
	// todo: add your logic here and delete this line

	return
}

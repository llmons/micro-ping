package blog

import (
	"context"

	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHotBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotBlogLogic {
	return &GetHotBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHotBlogLogic) GetHotBlog(req *types.ReqGetHotBlog) (resp *types.RespGetHotBlog, err error) {
	// todo: add your logic here and delete this line

	return
}

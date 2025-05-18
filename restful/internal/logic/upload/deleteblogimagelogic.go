package upload

import (
	"context"

	"restful/internal/svc"
	"restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBlogImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBlogImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBlogImageLogic {
	return &DeleteBlogImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBlogImageLogic) DeleteBlogImage(req *types.ReqDeleteBlogImage) (resp *types.RespDeleteBlogImage, err error) {
	// todo: add your logic here and delete this line

	return
}

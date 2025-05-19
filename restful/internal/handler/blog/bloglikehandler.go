package blog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"micro-ping/restful/internal/logic/blog"
	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"
)

func BlogLikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqLikeBlog
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := blog.NewBlogLikeLogic(r.Context(), svcCtx)
		resp, err := l.BlogLike(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

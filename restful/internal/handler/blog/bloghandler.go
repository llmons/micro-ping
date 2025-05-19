package blog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"micro-ping/restful/internal/logic/blog"
	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"
)

func BlogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqInsertBlog
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := blog.NewBlogLogic(r.Context(), svcCtx)
		resp, err := l.Blog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

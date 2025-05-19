package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"micro-ping/restful/internal/logic/user"
	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"
)

func GetInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqGetInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

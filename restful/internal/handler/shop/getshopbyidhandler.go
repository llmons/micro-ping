package shop

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"restful/internal/logic/shop"
	"restful/internal/svc"
	"restful/internal/types"
)

func GetShopByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqGetShopById
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := shop.NewGetShopByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetShopById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

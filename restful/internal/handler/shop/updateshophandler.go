package shop

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"micro-ping/restful/internal/logic/shop"
	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"
)

func UpdateShopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqUpdateShop
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := shop.NewUpdateShopLogic(r.Context(), svcCtx)
		resp, err := l.UpdateShop(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

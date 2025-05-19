package shop

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"micro-ping/restful/internal/logic/shop"
	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"
)

func InsertShopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqInsertShop
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := shop.NewInsertShopLogic(r.Context(), svcCtx)
		resp, err := l.InsertShop(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

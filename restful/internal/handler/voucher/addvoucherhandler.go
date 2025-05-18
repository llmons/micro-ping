package voucher

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"restful/internal/logic/voucher"
	"restful/internal/svc"
	"restful/internal/types"
)

func AddVoucherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqAddVoucher
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := voucher.NewAddVoucherLogic(r.Context(), svcCtx)
		resp, err := l.AddVoucher(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

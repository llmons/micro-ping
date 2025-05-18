package voucher_order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"restful/internal/logic/voucher_order"
	"restful/internal/svc"
	"restful/internal/types"
)

func SeckillVoucherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqSeckillVoucher
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := voucher_order.NewSeckillVoucherLogic(r.Context(), svcCtx)
		resp, err := l.SeckillVoucher(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

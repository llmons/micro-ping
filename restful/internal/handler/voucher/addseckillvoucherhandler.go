package voucher

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"micro-ping/restful/internal/logic/voucher"
	"micro-ping/restful/internal/svc"
	"micro-ping/restful/internal/types"
)

func AddSeckillVoucherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqAddSeckillVoucher
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := voucher.NewAddSeckillVoucherLogic(r.Context(), svcCtx)
		resp, err := l.AddSeckillVoucher(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

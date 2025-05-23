// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package types

type ReqAddSeckillVoucher struct {
	Voucher Voucher `json:"voucher"`
}

type ReqAddVoucher struct {
	Voucher Voucher `json:"voucher"`
}

type ReqGetVoucherByShopId struct {
	ShopId int64 `path:"shop_id"`
}

type RespAddSeckillVoucher struct {
	Id int64 `path:"id"`
}

type RespAddVoucher struct {
	Id int64 `path:"id"`
}

type RespGetVoucherByShopId struct {
	VoucherList []Voucher `json:"voucher_list"`
}

package response

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type Response struct {
	Success  bool        `json:"success"`
	Data     interface{} `json:"data,omitempty"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Total    int64       `json:"total,omitempty"`
}

func Success(data interface{}) *Response {
	return &Response{
		Success:  true,
		Data:     data,
		ErrorMsg: "",
		Total:    0,
	}
}

func Fail(err string) *Response {
	return &Response{
		Success:  false,
		Data:     nil,
		ErrorMsg: err,
		Total:    0,
	}
}

func OkHandler(_ context.Context, v interface{}) any {
	return Success(v)
}

func ErrHandler(name string) func(ctx context.Context, err error) (int, any) {
	return func(ctx context.Context, err error) (int, any) {
		logx.WithContext(ctx).Errorf("【%s】 err %v", name, err)
		return http.StatusBadRequest, Fail(err.Error())
	}
}

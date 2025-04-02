package result

import (
	"wrblog-api-go/pkg/constants"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// New 自定义返回
func New(code int, msg string, data any) *Result {
	return &Result{code, msg, data}
}

// Ok 成功返回
func Ok(data any) *Result {
	// 默认返回数据格式的结构体
	defaultData := &Result{
		Code: constants.SUCCESS,
		Msg:  "ok",
		Data: nil,
	}
	if data != nil {
		defaultData.Data = data
	}
	return &Result{Code: defaultData.Code, Msg: defaultData.Msg, Data: defaultData.Data}
}

// Fail 失败返回
func Fail(msg string) *Result {
	// 默认返回数据格式的结构体
	defaultData := &Result{
		Code: constants.ERROR,
		Msg:  "error",
		Data: nil,
	}
	if msg != "" {
		defaultData.Msg = msg
	}
	return &Result{Code: defaultData.Code, Msg: defaultData.Msg, Data: defaultData.Data}
}

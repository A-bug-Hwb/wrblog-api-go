package result

import (
	"wrblog-api-go/pkg/constants"
)

type Rows struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Rows  any    `json:"rows"`
	Total int64  `json:"total"`
}

// Suc 成功返回
func Suc(row any, total int64) *Rows {
	// 默认返回数据格式的结构体
	defaultData := &Rows{
		Code:  constants.SUCCESS,
		Msg:   "ok",
		Rows:  nil,
		Total: 0,
	}
	if row != nil {
		defaultData.Rows = row
	}
	if total != 0 {
		defaultData.Total = total
	}
	return &Rows{Code: defaultData.Code, Msg: defaultData.Msg, Rows: defaultData.Rows, Total: defaultData.Total}
}

// Err 失败返回
func Err(msg string) *Rows {
	// 默认返回数据格式的结构体
	defaultData := &Rows{
		Code:  constants.ERROR,
		Msg:   "error",
		Rows:  nil,
		Total: 0,
	}
	if msg != "" {
		defaultData.Msg = msg
	}
	return &Rows{Code: defaultData.Code, Msg: defaultData.Msg, Rows: defaultData.Rows, Total: defaultData.Total}
}

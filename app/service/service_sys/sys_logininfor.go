package service_sys

import (
	"fmt"
	"wrblog-api-go/app/dao/dao_sys"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/pkg/mylog"
)

// AddLoginifor 添加登录日志
func AddLoginifor(logininfor *model_sys.SysLogininfor) {
	err := dao_sys.AddLoginifor(logininfor)
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("登录日志添加失败！%s", err))
	}
}

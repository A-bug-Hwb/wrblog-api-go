package dao_sys

import (
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/pkg/mysql"
)

// AddLoginifor 添加登录日志
func AddLoginifor(logininfor *model_sys.SysLogininfor) (err error) {
	err = mysql.Db().Create(&logininfor).Error
	return
}

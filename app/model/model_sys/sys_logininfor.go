package model_sys

import (
	"wrblog-api-go/app/common/token"
	"wrblog-api-go/pkg/utils"
)

type SysLogininfor struct {
	InfoId           int        `json:"infoId" form:"infoId" gorm:"primaryKey;autoIncrement"` //日志id
	Ak               string     `json:"ak" form:"ak"`                                         //账号
	token.ClientInfo            //客户端信息
	Status           string     `json:"status" form:"status"`         //登录状态 0成功 1失败
	Msg              string     `json:"msg" form:"msg"`               //说明
	AccessTime       utils.Time `json:"accessTime" form:"accessTime"` //请求时间
}

func (SysLogininfor) TableName() string {
	return "sys_logininfor"
}

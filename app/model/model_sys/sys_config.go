package model_sys

import (
	"wrblog-api-go/app/model"
)

type SysConfig struct {
	ConfigId    uint   `json:"configId" form:"configId" gorm:"primaryKey;autoIncrement"`
	ConfigName  string `json:"configName" form:"configName"`
	ConfigKey   string `json:"configKey" form:"configKey"`
	ConfigValue string `json:"configValue" form:"configValue"`
	ConfigType  string `json:"configType" form:"configType"` //系统内置 Y是 N否
}

type SysConfigPo struct {
	SysConfig
	model.BaseEntity
}

func (SysConfigPo) TableName() string {
	return "sys_config"
}

type SysConfigVo struct {
	SysConfig
	model.BaseEntityVo
}

type SelectSysConfig struct {
	SysConfig
	model.SelectBaseEntity
}

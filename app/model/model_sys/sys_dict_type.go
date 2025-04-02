package model_sys

import "wrblog-api-go/app/model"

type SysDictType struct {
	DictId   int    `json:"dictId" form:"dictId" gorm:"primaryKey;autoIncrement"` //字典id
	DictName string `json:"dictName" form:"dictName" gorm:"required"`             //字典名称
	DictType string `json:"dictType" form:"dictType" gorm:"required"`             //字典类型
	Status   string `json:"status" form:"status" gorm:"required"`                 //状态
}

type SysDictTypePo struct {
	SysDictType
	model.BaseEntity
}

func (SysDictTypePo) TableName() string {
	return "sys_dict_type"
}

type SysDictTypeVo struct {
	SysDictType
	model.BaseEntityVo
}

type SelectSysDictType struct {
	SysDictType
	model.SelectBaseEntity
}

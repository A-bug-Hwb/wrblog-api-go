package model_sys

import (
	"encoding/json"
	"wrblog-api-go/app/model"
)

type SysDictData struct {
	DictCode  int    `json:"dictCode" form:"dictId" gorm:"primaryKey;autoIncrement"` //字典编码
	DictSort  int    `json:"dictSort" form:"dictSort"`                               //字典排序
	DictLabel string `json:"dictLabel" form:"dictLabel" gorm:"required"`             //字典标签
	DictValue string `json:"dictValue" form:"dictValue" gorm:"required"`             //字典键值
	DictType  string `json:"dictType" form:"dictType" gorm:"required"`               //字典类型
	CssClass  string `json:"cssClass" form:"cssClass"`                               //样式属性（其他样式扩展）
	ListClass string `json:"listClass" form:"listClass"`                             //表格回显样式
	IsDefault string `json:"isDefault" form:"isDefault"`                             //是否默认（Y是 N否）
	Status    string `json:"status" form:"status" gorm:"required"`                   //状态（0正常 1停用）
}

type SysDictDataPo struct {
	SysDictData
	model.BaseEntity
}

func (SysDictDataPo) TableName() string {
	return "sys_dict_data"
}

type SysDictDataVo struct {
	SysDictData
	model.BaseEntityVo
}

type SelectSysDictData struct {
	SysDictData
	model.SelectBaseEntity
}

// 序列化
func (sysDictData *SysDictDataVo) MarshalBinary() (data []byte, err error) {
	return json.Marshal(sysDictData)
}

// 反序列化
func (sysDictData *SysDictDataVo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, sysDictData)
}

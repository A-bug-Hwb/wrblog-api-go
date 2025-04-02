package model

import (
	"wrblog-api-go/pkg/utils"
)

type SelectBaseEntity struct {
	PageNum   int        `json:"pageNum" form:"pageNum" gorm:"-"`     //当前页（非分页接口不传）
	PageSize  int        `json:"pageSize" form:"pageSize" gorm:"-"`   //每页条数（非分页接口不传）
	StartTime utils.Time `json:"startTime" form:"startTime" gorm:"-"` //开始时间
	EndTime   utils.Time `json:"endTime" form:"endTime" gorm:"-"`     //开始时间
}

type BaseEntityVo struct {
	CreateTime utils.Time             `json:"createTime" form:"createTime"`  //创建时间
	Remark     string                 `json:"remark" form:"remark"`          //备注
	Params     map[string]interface{} `json:"params" form:"params" gorm:"-"` //拓展字段
}

type BaseEntity struct {
	CreateBy   string     `json:"createBy" form:"createBy"`                           //创建者
	CreateTime utils.Time `json:"createTime" form:"createTime" gorm:"autoCreateTime"` //创建时间
	UpdateBy   string     `json:"updateBy" form:"updateBy"`                           //修改者
	UpdateTime utils.Time `json:"updateTime" form:"updateTime" gorm:"autoUpdateTime"` //修改时间
	Remark     string     `json:"remark" form:"remark"`                               //备注
	Deleted    string     `json:"deleted" form:"deleted" gorm:"default:0"`            //逻辑删除
}

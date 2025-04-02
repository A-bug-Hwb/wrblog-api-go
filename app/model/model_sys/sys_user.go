package model_sys

import (
	"wrblog-api-go/app/model"
	"wrblog-api-go/pkg/utils"
)

type SysUser struct {
	DeptId       int        `json:"deptId" form:"deptId"`
	Uk           string     `json:"uk" form:"uk"`
	NickName     string     `json:"nickName" form:"nickName"`
	SpaceUrl     string     `json:"spaceUrl" form:"spaceUrl"`
	Mobile       string     `json:"mobile" form:"mobile"`
	Mailbox      string     `json:"mailbox" form:"mailbox"`
	Avatar       string     `json:"avatar" form:"avatar"`
	Sex          string     `json:"sex" form:"sex"`
	RegisterType string     `json:"registerType" form:"registerType"`
	UserType     string     `json:"userType" form:"userType"`
	LoginIp      string     `json:"loginIp" form:"loginIp"`
	LoginDate    utils.Time `json:"loginDate" form:"loginDate"`
	Status       string     `json:"status" form:"status"`
}

type SysUserPo struct {
	UserId int `json:"userId" form:"userId" gorm:"primaryKey;autoIncrement"`
	SysUser
	Pk string `json:"pk" form:"pk"`
	model.BaseEntity
}

func (SysUserPo) TableName() string {
	return "sys_user"
}

type SysUserVo struct {
	UserId string `json:"userId" form:"userId"`
	SysUser
	model.BaseEntityVo
}

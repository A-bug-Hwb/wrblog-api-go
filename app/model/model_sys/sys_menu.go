package model_sys

import (
	"wrblog-api-go/app/model"
)

type SysMenu struct {
	MenuId    int    `json:"menuId" form:"userId" gorm:"primaryKey;autoIncrement"` //菜单id
	MenuName  string `json:"menuName" form:"menuName" gorm:"required"`             //菜单名称
	ParentId  int    `json:"parentId" form:"parentId"`                             //上级id
	OrderNum  int    `json:"orderNum" form:"orderNum"`                             //排序
	Path      string `json:"path" form:"path" gorm:"required"`                     //路由地址
	Component string `json:"component" form:"component"`                           //组件路径
	Query     string `json:"query" form:"query"`                                   //参数
	IsFrame   string `json:"isFrame" form:"isFrame"`                               //是否外链 0 是 1否
	MenuType  string `json:"menuType" form:"menuType" gorm:"required"`             //菜单类型
	Visible   string `json:"visible" form:"visible"`                               //菜单状态（0显示 1隐藏）
	Status    string `json:"status" form:"status"`                                 //菜单状态（0正常 1停用）
	Perms     string `json:"perms" form:"perms"`                                   //权限标识
	Icon      string `json:"icon" form:"icon"`                                     //图标
}

type SysMenuPo struct {
	SysMenu
	model.BaseEntity
}

func (SysMenuPo) TableName() string {
	return "sys_menu"
}

type SelectSysMenu struct {
	MenuName string `json:"menuName" form:"menuName"` //菜单名
	Status   string `json:"status" form:"status"`     //状态
	UserId   int    `json:"userId" form:"userId"`     //用户id
	RoleId   int    `json:"roleId" form:"roleId"`     //角色id
	model.SelectBaseEntity
}

type SysMenuVo struct {
	SysMenu
	Children []*SysMenuVo `json:"children" gorm:"-"`
	model.BaseEntityVo
}

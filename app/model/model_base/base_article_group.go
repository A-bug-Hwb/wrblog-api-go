package model_base

import "wrblog-api-go/app/model"

type BaseArticleGroup struct {
	GroupId   int    `json:"groupId" form:"groupId" gorm:"primaryKey;autoIncrement"` //组id
	GroupName string `json:"groupName" form:"groupName"`                             //组名
	UserId    int    `json:"userId" form:"userId"`                                   //用户id
	LibraryId int    `json:"libraryId" form:"libraryId"`                             //库id
	Status    string `json:"status" form:"status"`                                   //状态
}

type BaseArticleGroupPo struct {
	BaseArticleGroup
	model.BaseEntity
}

func (BaseArticleGroupPo) TableName() string {
	return "base_article_group"
}

type SelectBaseArticleGroup struct {
	BaseArticleGroup
	model.SelectBaseEntity
}

type BaseArticleGroupVo struct {
	BaseArticleGroup
	model.BaseEntityVo
}

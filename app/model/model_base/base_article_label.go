package model_base

import "wrblog-api-go/app/model"

type BaseArticleLabel struct {
	LabelId   int    `json:"labelId" form:"labelId" gorm:"primaryKey;autoIncrement"` //标签id
	LabelName string `json:"labelName" form:"labelName"`                             //标签名
	UserId    int    `json:"userId" form:"userId"`                                   //用户id
	Status    string `json:"status" form:"status"`                                   //状态
}

type BaseArticleLabelPo struct {
	BaseArticleLabel
	model.BaseEntity
}

func (BaseArticleLabelPo) TableName() string {
	return "base_article_label"
}

type SelectBaseArticleLabel struct {
	BaseArticleLabel
	model.SelectBaseEntity
}

type BaseArticleLabelVo struct {
	BaseArticleLabel
	model.BaseEntityVo
}

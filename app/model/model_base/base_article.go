package model_base

import "wrblog-api-go/app/model"

type BaseArticle struct {
	ArticleId  int    `json:"articleId" form:"articleId" gorm:"primaryKey;autoIncrement"` //文章id
	ParentId   int    `json:"parentId" form:"parentId"`                                   //父级id
	LibraryId  int    `json:"libraryId" form:"libraryId"`                                 //库id
	ArticleUrl string `json:"articleUrl" form:"articleUrl"`                               //文章路径
	UserId     int    `json:"userId" form:"userId"`                                       //用户id
	Title      string `json:"title" form:"title" gorm:"required"`                         //标题
	Content    string `json:"content" form:"content" gorm:"required"`                     //内容
	LabelIds   string `json:"labelIds" form:"labelIds"`                                   //标签
	TypeId     int    `json:"typeId" form:"typeId"`                                       //类型
	GroupId    int    `json:"groupId" form:"groupId"`                                     //组id
	IsOpen     string `json:"IsOpen" form:"IsOpen"`                                       //是否公开文章
	Status     string `json:"status" form:"status"`                                       //状态
}

type BaseArticlePo struct {
	BaseArticle
	model.BaseEntity
}

func (BaseArticlePo) TableName() string {
	return "base_article"
}

type SelectBaseArticle struct {
	BaseArticle
	model.SelectBaseEntity
}

type BaseArticleVo struct {
	BaseArticle
	model.BaseEntityVo
}

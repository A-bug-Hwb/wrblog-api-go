package model_base

import "wrblog-api-go/app/model"

type BaseArticleLibrary struct {
	LibraryId   int    `json:"libraryId" form:"libraryId" gorm:"primaryKey;autoIncrement"` //库id
	LibraryName string `json:"libraryName" form:"libraryName"`                             //库名
	LibraryUrl  string `json:"libraryUrl" form:"libraryUrl"`                               //库的路径
	UserId      int    `json:"userId" form:"userId"`                                       //用户id
	IsOpen      string `json:"IsOpen" form:"IsOpen"`                                       //是否公开库
	Status      string `json:"status" form:"status"`                                       //状态
}

type BaseArticleLibraryPo struct {
	BaseArticleLibrary
	model.BaseEntity
}

func (BaseArticleLibraryPo) TableName() string {
	return "base_article_library"
}

type SelectBaseArticleLibrary struct {
	BaseArticleLibrary
	model.SelectBaseEntity
}

type BaseArticleLibraryVo struct {
	BaseArticleLibrary
	model.BaseEntityVo
}

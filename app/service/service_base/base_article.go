package service_base

import (
	"wrblog-api-go/app/dao/dao_base"
	"wrblog-api-go/app/model/model_base"
)

func GetArticlePageList(baseArticle *model_base.SelectBaseArticle) ([]*model_base.BaseArticleVo, int64) {
	return dao_base.GetArticlePageList(baseArticle)
}

func GetArticleById(articleId int) *model_base.BaseArticleVo {
	return dao_base.GetArticleById(articleId)
}

func AddArticle(baseArticle *model_base.BaseArticlePo) int {
	_, articleId := dao_base.SaveArticle(baseArticle)
	return articleId
}

func EditArticle(baseArticle *model_base.BaseArticlePo) int64 {
	row, _ := dao_base.SaveArticle(baseArticle)
	return row
}

func RemoveArticle(articleIds []string, uk string) int64 {
	return dao_base.RemoveArticle(articleIds, uk)
}

func DeleteArticle(articleIds []string) int64 {
	return dao_base.DeleteArticle(articleIds)
}

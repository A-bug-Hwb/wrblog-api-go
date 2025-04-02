package service_base

import (
	"wrblog-api-go/app/dao/dao_base"
	"wrblog-api-go/app/model/model_base"
)

func GetArticleLabelPageList(baseArticleLabel *model_base.SelectBaseArticleLabel) ([]*model_base.BaseArticleLabelVo, int64) {
	return dao_base.GetArticleLabelPageList(baseArticleLabel)
}

func GetArticleLabelInfo(labelId int) *model_base.BaseArticleLabelVo {
	return dao_base.GetArticleLabelInfo(labelId)
}

func AddArticleLabel(baseArticleLabel *model_base.BaseArticleLabelPo) int {
	_, labelId := dao_base.SaveArticleLabel(baseArticleLabel)
	return labelId
}

func EditArticleLabel(baseArticleLabel *model_base.BaseArticleLabelPo) int64 {
	row, _ := dao_base.SaveArticleLabel(baseArticleLabel)
	return row
}

// RemoveArticleLabel 逻辑删除
func RemoveArticleLabel(labelIds []string, uk string) int64 {
	return dao_base.RemoveArticleLabel(labelIds, uk)
}

// DeletedArticleLabel 删除
func DeletedArticleLabel(labelIds []string) int64 {
	return dao_base.DeletedArticleLabel(labelIds)
}

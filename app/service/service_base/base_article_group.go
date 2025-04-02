package service_base

import (
	"wrblog-api-go/app/dao/dao_base"
	"wrblog-api-go/app/model/model_base"
)

func GetArticleGroupPageList(baseArticleGroup *model_base.SelectBaseArticleGroup) ([]*model_base.BaseArticleGroupVo, int64) {
	return dao_base.GetArticleGroupPageList(baseArticleGroup)
}

func GetArticleGroupInfo(groupId int) (baseArticleGroup *model_base.BaseArticleGroupVo) {
	return dao_base.GetArticleGroupInfo(groupId)
}

func AddArticleGroup(baseArticleGroup *model_base.BaseArticleGroupPo) int {
	_, groupId := dao_base.SaveArticleGroup(baseArticleGroup)
	return groupId
}

func EditArticleGroup(baseArticleGroup *model_base.BaseArticleGroupPo) int64 {
	row, _ := dao_base.SaveArticleGroup(baseArticleGroup)
	return row
}

// RemoveArticleGroup 逻辑删除
func RemoveArticleGroup(groupIds []string, uk string) int64 {
	return dao_base.RemoveArticleGroup(groupIds, uk)
}

// DeletedArticleGroup 删除
func DeletedArticleGroup(groupIds []string) int64 {
	return dao_base.DeletedArticleGroup(groupIds)
}

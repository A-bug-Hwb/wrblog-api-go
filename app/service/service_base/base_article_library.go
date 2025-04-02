package service_base

import (
	"wrblog-api-go/app/dao/dao_base"
	"wrblog-api-go/app/model/model_base"
)

func GetArticleLibraryPageList(baseArticleLibrary *model_base.SelectBaseArticleLibrary) ([]*model_base.BaseArticleLibraryVo, int64) {
	return dao_base.GetArticleLibraryPageList(baseArticleLibrary)
}

func GetArticleLibraryInfo(libraryId int) *model_base.BaseArticleLibraryVo {
	return dao_base.GetArticleLibraryInfo(libraryId)
}

func AddArticleLibrary(baseArticleLibrary *model_base.BaseArticleLibraryPo) int {
	_, libraryId := dao_base.SaveArticleLibrary(baseArticleLibrary)
	return libraryId
}

func EditArticleLibrary(baseArticleLibrary *model_base.BaseArticleLibraryPo) int64 {
	row, _ := dao_base.SaveArticleLibrary(baseArticleLibrary)
	return row
}

// RemoveArticleLibrary 逻辑删除
func RemoveArticleLibrary(libraryIds []string, uk string) int64 {
	return dao_base.RemoveArticleLibrary(libraryIds, uk)
}

// DeletedArticleLibrary 删除
func DeletedArticleLibrary(libraryIds []string) int64 {
	return dao_base.DeletedArticleLibrary(libraryIds)
}

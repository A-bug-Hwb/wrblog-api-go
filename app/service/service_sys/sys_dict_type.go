package service_sys

import (
	"wrblog-api-go/app/dao/dao_sys"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/pkg/mylog"
)

func GetDictTypePageList(sysDictType *model_sys.SelectSysDictType) ([]*model_sys.SysDictTypeVo, int64) {
	return dao_sys.GetDictTypePageList(sysDictType)
}

func GetDictTypeById(dictId int) *model_sys.SysDictTypeVo {
	return dao_sys.GetDictTypeById(dictId)
}

func AddDictType(sysDictType *model_sys.SysDictTypePo) int {
	_, dictId := dao_sys.SaveDictType(sysDictType)
	return dictId
}
func EditDictType(sysDictType *model_sys.SysDictTypePo) int64 {
	if sysDictType.DictId == 0 {
		mylog.MyLog.Panic("缺少参数：dictId")
	}
	row, _ := dao_sys.SaveDictType(sysDictType)
	return row
}

func RemoveDictType(dictIds []string, uk string) int64 {
	return dao_sys.RemoveDictType(dictIds, uk)
}

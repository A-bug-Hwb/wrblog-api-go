package service_sys

import (
	"encoding/json"
	"fmt"
	"wrblog-api-go/app/common/constants"
	"wrblog-api-go/app/dao/dao_sys"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/redis"
)

func GetDictDataPageList(sysDictData *model_sys.SelectSysDictData) ([]*model_sys.SysDictDataVo, int64) {
	return dao_sys.GetDictDataPageList(sysDictData)
}
func ApiDictDataByType(dictType string) (sysDictDatas []*model_sys.SysDictDataVo) {
	val, _ := redis.Get(fmt.Sprintf("%s%s", constants.SYS_DICT, dictType))
	if val != nil {
		json.Unmarshal(val, &sysDictDatas)
	} else {
		sysDictDatas = dao_sys.ApiDictDataByType(dictType)
		info, _ := json.Marshal(sysDictDatas)
		redis.Set(fmt.Sprintf("%s%s", constants.SYS_DICT, dictType), info)
	}
	return
}

func GetDictDataById(dictCode int) *model_sys.SysDictDataVo {
	return dao_sys.GetDictDataById(dictCode)
}

func AddDictData(sysDictData *model_sys.SysDictDataPo) int {
	_, dictCode := dao_sys.SaveDictData(sysDictData)
	return dictCode
}

func EditDictData(sysDictData *model_sys.SysDictDataPo) int64 {
	if sysDictData.DictCode == 0 {
		mylog.MyLog.Panic("缺少参数：dictCode")
	}
	row, _ := dao_sys.SaveDictData(sysDictData)
	return row
}

func RemoveDictData(dictIds []string, uk string) int64 {
	return dao_sys.RemoveDictData(dictIds, uk)
}

func RemoveCacheAll() int {
	keys, _ := redis.ScanKeys(fmt.Sprintf("%s%s", constants.SYS_DICT, "*"))
	redis.Del(keys...)
	return len(keys)
}

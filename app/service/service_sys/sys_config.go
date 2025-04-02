package service_sys

import (
	"wrblog-api-go/app/common/constants"
	"wrblog-api-go/app/dao/dao_sys"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/pkg/redis"
)

func init() {
	//CacheConfigList()
}

// CacheConfigList 缓存参数列表
func CacheConfigList() {
	sysConfigs, _ := GetPageList(&model_sys.SelectSysConfig{})
	for _, sysConfig := range sysConfigs {
		redis.Set(constants.SYS_CONFIG+sysConfig.ConfigKey, sysConfig.ConfigValue)
	}
}

func GetPageList(sysConfig *model_sys.SelectSysConfig) ([]*model_sys.SysConfigVo, int64) {
	return dao_sys.GetPageList(sysConfig)
}

func GetInfoByKey(key string) *model_sys.SysConfigPo {
	return dao_sys.GetInfoByKey(key)
}

func GetValueByKey(key string) string {
	return dao_sys.GetValueByKey(key)
}

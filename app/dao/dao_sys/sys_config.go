package dao_sys

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"wrblog-api-go/app/common/constants"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/mysql"
	"wrblog-api-go/pkg/redis"
)

func GetPageList(sysConfig *model_sys.SelectSysConfig) (sysConfigs []*model_sys.SysConfigVo, total int64) {
	db := mysql.Db().Model(&model_sys.SysConfigPo{})
	if sysConfig.ConfigName != "" {
		db.Where("config_name like ?", "%"+sysConfig.ConfigName+"%")
	}
	if sysConfig.ConfigKey != "" {
		db.Where("config_key like ?", "%"+sysConfig.ConfigKey+"%")
	}
	if !sysConfig.StartTime.IsZero() && !sysConfig.EndTime.IsZero() {
		db.Where("create_time BETWEEN ? AND ?", sysConfig.StartTime, sysConfig.EndTime)
	}
	err := db.Where("deleted = 0").Count(&total).Find(&sysConfigs).Error
	if sysConfig.PageNum != 0 && sysConfig.PageSize != 0 {
		db.Limit(sysConfig.PageSize).Offset((sysConfig.PageNum - 1) * sysConfig.PageSize)
	} else {
		db.Limit(10)
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetInfoByKey(key string) (sysConfig *model_sys.SysConfigPo) {
	err := mysql.Db().Where("config_key = ? and deleted = 0", key).First(&sysConfig).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetValueByKey(key string) (val string) {
	byteVal, err := redis.Get(constants.SYS_CONFIG + key)
	if err != nil {
		var sysConfig model_sys.SysConfigPo
		err = mysql.Db().Where("config_key = ? and deleted = 0", key).First(&sysConfig).Error
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
			}
			return ""
		}
		redis.Set(constants.SYS_CONFIG+key, sysConfig.ConfigValue)
		return sysConfig.ConfigValue
	} else {
		return string(byteVal)
	}
}

package dao_sys

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"wrblog-api-go/app/model"
	"wrblog-api-go/app/model/model_base"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/mysql"
	"wrblog-api-go/pkg/utils"
)

func GetDictTypePageList(sysDictType *model_sys.SelectSysDictType) (sysDictTypes []*model_sys.SysDictTypeVo, total int64) {
	db := mysql.Db().Model(&model_sys.SysDictTypePo{})
	if sysDictType.DictName != "" {
		db.Where("dict_name like ?", "%"+sysDictType.DictName+"%")
	}
	if sysDictType.DictType != "" {
		db.Where("dict_type like ?", "%"+sysDictType.DictType+"%")
	}
	if sysDictType.Status != "" {
		db.Where("status = ?", sysDictType.Status)
	}
	if !sysDictType.StartTime.IsZero() && !sysDictType.EndTime.IsZero() {
		db.Where("create_time BETWEEN ? AND ?", sysDictType.StartTime, sysDictType.EndTime)
	}
	db.Where("deleted = 0").Count(&total).Order("create_time desc")
	if sysDictType.PageNum != 0 && sysDictType.PageSize != 0 {
		db.Limit(sysDictType.PageSize).Offset((sysDictType.PageNum - 1) * sysDictType.PageSize)
	} else {
		db.Limit(10)
	}
	err := db.Find(&sysDictTypes).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetDictTypeById(dictId int) (sysDictType *model_sys.SysDictTypeVo) {
	db := mysql.Db().Model(&model_sys.SysDictTypePo{})
	err := db.Where("dict_id = ? and deleted = 0", dictId).First(&sysDictType).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func SaveDictType(sysDictType *model_sys.SysDictTypePo) (int64, int) {
	res := mysql.Db().Save(&sysDictType)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected, sysDictType.DictId
}

func RemoveDictType(dictIds []string, uk string) int64 {
	db := mysql.Db().Model(&model_sys.SysDictTypePo{})
	db.Where("dict_id in (?)", dictIds)
	res := db.Updates(&model_sys.SysDictTypePo{BaseEntity: model.BaseEntity{UpdateTime: utils.Time(time.Now()), UpdateBy: uk, Deleted: "1"}})
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

func DeleteDictType(dictIds []string) int64 {
	db := mysql.Db().Model(&model_base.BaseArticlePo{})
	res := db.Delete(&model_base.BaseArticlePo{}, dictIds)
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

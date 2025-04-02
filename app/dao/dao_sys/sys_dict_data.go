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

func GetDictDataPageList(sysDictData *model_sys.SelectSysDictData) (sysDictDatas []*model_sys.SysDictDataVo, total int64) {
	db := mysql.Db().Model(&model_sys.SysDictDataPo{})
	if sysDictData.DictLabel != "" {
		db.Where("dict_label like ?", "%"+sysDictData.DictLabel+"%")
	}
	if sysDictData.DictType != "" {
		db.Where("dict_type = ?", sysDictData.DictType)
	}
	if sysDictData.Status != "" {
		db.Where("status = ?", sysDictData.Status)
	}
	if !sysDictData.StartTime.IsZero() && !sysDictData.EndTime.IsZero() {
		db.Where("create_time BETWEEN ? AND ?", sysDictData.StartTime, sysDictData.EndTime)
	}
	db.Where("deleted = 0").Count(&total).Order("dict_sort")
	if sysDictData.PageNum != 0 && sysDictData.PageSize != 0 {
		db.Limit(sysDictData.PageSize).Offset((sysDictData.PageNum - 1) * sysDictData.PageSize)
	} else {
		db.Limit(10)
	}
	err := db.Find(&sysDictDatas).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func ApiDictDataByType(dictType string) (sysDictDatas []*model_sys.SysDictDataVo) {
	db := mysql.Db().Model(&model_sys.SysDictDataPo{})
	db.Where("deleted = 0 and dict_type = ?", dictType).Order("dict_sort")
	err := db.Find(&sysDictDatas).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetDictDataById(dictCode int) (sysDictData *model_sys.SysDictDataVo) {
	db := mysql.Db().Model(&model_sys.SysDictDataPo{})
	err := db.Where("dict_code = ? and deleted = 0", dictCode).First(&sysDictData).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func SaveDictData(sysDictData *model_sys.SysDictDataPo) (int64, int) {
	res := mysql.Db().Save(&sysDictData)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected, sysDictData.DictCode
}

func RemoveDictData(dictCodes []string, uk string) int64 {
	db := mysql.Db().Model(&model_sys.SysDictDataPo{})
	db.Where("dict_code in (?)", dictCodes)
	res := db.Updates(&model_sys.SysDictDataPo{BaseEntity: model.BaseEntity{UpdateTime: utils.Time(time.Now()), UpdateBy: uk, Deleted: "1"}})
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

func DeleteDictData(dictIds []string) int64 {
	db := mysql.Db().Model(&model_base.BaseArticlePo{})
	res := db.Delete(&model_base.BaseArticlePo{}, dictIds)
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

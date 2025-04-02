package dao_base

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"wrblog-api-go/app/model"
	"wrblog-api-go/app/model/model_base"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/mysql"
	"wrblog-api-go/pkg/utils"
)

func GetArticleLibraryPageList(baseArticleLibrary *model_base.SelectBaseArticleLibrary) (baseArticleLibrarys []*model_base.BaseArticleLibraryVo, total int64) {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticleLibraryPo{})
	if baseArticleLibrary.UserId != 0 {
		db.Where("user_id = ?", baseArticleLibrary.UserId)
	}
	if baseArticleLibrary.LibraryName != "" {
		db.Where("library_name like ?", "%"+baseArticleLibrary.LibraryName+"%")
	}
	if baseArticleLibrary.LibraryUrl != "" {
		db.Where("library_url like ?", "%"+baseArticleLibrary.LibraryUrl+"%")
	}
	if baseArticleLibrary.IsOpen != "" {
		db.Where("is_open = ?", baseArticleLibrary.IsOpen)
	}
	if baseArticleLibrary.Status != "" {
		db.Where("status = ?", baseArticleLibrary.Status)
	}
	if !baseArticleLibrary.StartTime.IsZero() && !baseArticleLibrary.EndTime.IsZero() {
		db.Where("create_time BETWEEN ? AND ?", baseArticleLibrary.StartTime, baseArticleLibrary.EndTime)
	}
	db.Where("deleted = 0").Count(&total).Order("create_time desc")
	if baseArticleLibrary.PageNum != 0 && baseArticleLibrary.PageSize != 0 {
		db.Limit(baseArticleLibrary.PageSize).Offset((baseArticleLibrary.PageNum - 1) * baseArticleLibrary.PageSize)
	} else {
		db.Limit(10)
	}
	err := db.Find(&baseArticleLibrarys).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetArticleLibraryInfo(libraryId int) (baseArticleLibrary *model_base.BaseArticleLibraryVo) {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticleLibraryPo{})
	err := db.Where("library_id = ? and deleted = 0", libraryId).First(&baseArticleLibrary).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func SaveArticleLibrary(baseArticleLibrary *model_base.BaseArticleLibraryPo) (int64, int) {
	db := mysql.Db("wr-base")
	res := db.Save(&baseArticleLibrary)
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected, baseArticleLibrary.LibraryId
}

// RemoveArticleLibrary 逻辑删除
func RemoveArticleLibrary(libraryIds []string, uk string) int64 {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticleLibraryPo{})
	db.Where("library_id in (?)", libraryIds)
	res := db.Updates(&model_base.BaseArticleLibraryPo{BaseEntity: model.BaseEntity{UpdateTime: utils.Time(time.Now()), UpdateBy: uk, Deleted: "1"}})
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

// DeletedArticleLibrary 删除
func DeletedArticleLibrary(libraryIds []string) int64 {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticleLibraryPo{})
	db.Where("library_id in (?)", libraryIds)
	res := db.Delete(&model_base.BaseArticleLibraryPo{})
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

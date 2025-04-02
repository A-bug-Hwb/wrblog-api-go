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

func GetArticleGroupPageList(baseArticleGroup *model_base.SelectBaseArticleGroup) (baseArticleGroups []*model_base.BaseArticleGroupVo, total int64) {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticleGroupPo{})
	if baseArticleGroup.UserId != 0 {
		db.Where("user_id = ?", baseArticleGroup.UserId)
	}
	if baseArticleGroup.GroupName != "" {
		db.Where("group_name like ?", "%"+baseArticleGroup.GroupName+"%")
	}
	if baseArticleGroup.LibraryId != 0 {
		db.Where("library_id = ?", baseArticleGroup.LibraryId)
	}
	if baseArticleGroup.Status != "" {
		db.Where("status = ?", baseArticleGroup.Status)
	}
	if !baseArticleGroup.StartTime.IsZero() && !baseArticleGroup.EndTime.IsZero() {
		db.Where("create_time BETWEEN ? AND ?", baseArticleGroup.StartTime, baseArticleGroup.EndTime)
	}
	db.Where("deleted = 0").Count(&total).Order("create_time desc")
	if baseArticleGroup.PageNum != 0 && baseArticleGroup.PageSize != 0 {
		db.Limit(baseArticleGroup.PageSize).Offset((baseArticleGroup.PageNum - 1) * baseArticleGroup.PageSize)
	} else {
		db.Limit(10)
	}
	err := db.Find(&baseArticleGroups).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetArticleGroupInfo(groupId int) (baseArticleGroup *model_base.BaseArticleGroupVo) {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticleGroupPo{})
	err := db.Where("group_id = ? and deleted = 0", groupId).First(&baseArticleGroup).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func SaveArticleGroup(baseArticleGroup *model_base.BaseArticleGroupPo) (int64, int) {
	db := mysql.Db("wr-base")
	res := db.Save(&baseArticleGroup)
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected, baseArticleGroup.GroupId
}

func RemoveArticleGroup(groupIds []string, uk string) int64 {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticleGroupPo{})
	db.Where("group_id in (?)", groupIds)
	res := db.Updates(&model_base.BaseArticleGroupPo{BaseEntity: model.BaseEntity{UpdateTime: utils.Time(time.Now()), UpdateBy: uk, Deleted: "1"}})
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

func DeletedArticleGroup(groupIds []string) int64 {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticleGroupPo{})
	db.Where("group_id in (?)", groupIds)
	res := db.Delete(&model_base.BaseArticleGroupPo{})
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

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

func GetArticleLabelPageList(baseArticleLabel *model_base.SelectBaseArticleLabel) (baseArticleLabels []*model_base.BaseArticleLabelVo, total int64) {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticleLabelPo{})
	if baseArticleLabel.UserId != 0 {
		db.Where("user_id = ?", baseArticleLabel.UserId)
	}
	if baseArticleLabel.LabelName != "" {
		db.Where("label_name like ?", "%"+baseArticleLabel.LabelName+"%")
	}
	if baseArticleLabel.Status != "" {
		db.Where("status = ?", baseArticleLabel.Status)
	}
	if !baseArticleLabel.StartTime.IsZero() && !baseArticleLabel.EndTime.IsZero() {
		db.Where("create_time BETWEEN ? AND ?", baseArticleLabel.StartTime, baseArticleLabel.EndTime)
	}
	db.Where("deleted = 0").Count(&total).Order("create_time desc")
	if baseArticleLabel.PageNum != 0 && baseArticleLabel.PageSize != 0 {
		db.Limit(baseArticleLabel.PageSize).Offset((baseArticleLabel.PageNum - 1) * baseArticleLabel.PageSize)
	} else {
		db.Limit(10)
	}
	err := db.Find(&baseArticleLabels).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetArticleLabelInfo(labelId int) (baseArticleLabel *model_base.BaseArticleLabelVo) {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticleLabelPo{})
	err := db.Where("label_id = ? and deleted = 0", labelId).First(&baseArticleLabel).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func SaveArticleLabel(baseArticleLabel *model_base.BaseArticleLabelPo) (int64, int) {
	db := mysql.Db("wr-base")
	res := db.Save(&baseArticleLabel)
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected, baseArticleLabel.LabelId
}

func RemoveArticleLabel(labelIds []string, uk string) int64 {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticleLabelPo{})
	db.Where("label_id in (?)", labelIds)
	res := db.Updates(&model_base.BaseArticleLabelPo{BaseEntity: model.BaseEntity{UpdateTime: utils.Time(time.Now()), UpdateBy: uk, Deleted: "1"}})
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

func DeletedArticleLabel(labelIds []string) int64 {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticleLabelPo{})
	db.Where("label_id in (?)", labelIds)
	res := db.Delete(&model_base.BaseArticleLabelPo{})
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

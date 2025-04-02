package dao_base

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"wrblog-api-go/app/model"
	"wrblog-api-go/app/model/model_base"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/mysql"
	"wrblog-api-go/pkg/utils"
)

func GetArticlePageList(baseArticle *model_base.SelectBaseArticle) (baseArticles []*model_base.BaseArticleVo, total int64) {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticlePo{})
	if baseArticle.Title != "" {
		db.Where("title like ?", "%"+baseArticle.Title+"%")
	}
	if baseArticle.UserId != 0 {
		db.Where("user_id = ?", baseArticle.UserId)
	}
	if baseArticle.LibraryId != 0 {
		db.Where("library_id = ?", baseArticle.LibraryId)
	}
	if baseArticle.LabelIds != "" {
		labelIds := strings.Split(baseArticle.LabelIds, ",")
		db.Where("label_ids in (?)", labelIds)
	}
	if baseArticle.TypeId != 0 {
		db.Where("type_id = ?", baseArticle.TypeId)
	}
	if baseArticle.GroupId != 0 {
		db.Where("group_id = ?", baseArticle.GroupId)
	}
	if baseArticle.IsOpen != "" {
		db.Where("is_open = ?", baseArticle.IsOpen)
	}
	if baseArticle.Status != "" {
		db.Where("status = ?", baseArticle.Status)
	}
	if !baseArticle.StartTime.IsZero() && !baseArticle.EndTime.IsZero() {
		db.Where("create_time BETWEEN ? AND ?", baseArticle.StartTime, baseArticle.EndTime)
	}
	db.Where("deleted = 0").Count(&total).Order("create_time desc")
	if baseArticle.PageNum != 0 && baseArticle.PageSize != 0 {
		db.Limit(baseArticle.PageSize).Offset((baseArticle.PageNum - 1) * baseArticle.PageSize)
	} else {
		db.Limit(10)
	}
	err := db.Find(&baseArticles).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetArticleById(articleId int) (baseArticle *model_base.BaseArticleVo) {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticlePo{})
	err := db.Where("article_id = ? and deleted = 0", articleId).First(&baseArticle).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func SaveArticle(baseArticle *model_base.BaseArticlePo) (int64, int) {
	db := mysql.Db("wr-base")
	res := db.Save(&baseArticle)
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected, baseArticle.ArticleId
}

func RemoveArticle(articleIds []string, uk string) int64 {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticlePo{})
	db.Where("article_id in (?)", articleIds)
	res := db.Updates(&model_base.BaseArticlePo{BaseEntity: model.BaseEntity{UpdateTime: utils.Time{}, UpdateBy: uk, Deleted: "1"}})
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

func DeleteArticle(articleIds []string) int64 {
	db := mysql.Db("wr-base").Model(&model_base.BaseArticlePo{})
	db.Where("article_id in (?)", articleIds)
	res := db.Delete(&model_base.BaseArticlePo{})
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

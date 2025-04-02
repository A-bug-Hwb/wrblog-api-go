package dao_sys

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"wrblog-api-go/app/model"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/mysql"
	"wrblog-api-go/pkg/utils"
)

func GetMenuList(selectSysMenu *model_sys.SelectSysMenu) (sysMenuList []*model_sys.SysMenuVo) {
	db := mysql.Db().Model(&model_sys.SysMenuPo{}).Table("sys_menu sm").Select("sm.*")
	db.Joins("left join sys_role_menu srm on srm.menu_id = sm.menu_id")
	db.Joins("left join sys_user_role sur on sur.role_id = srm.role_id")
	if selectSysMenu.MenuName != "" {
		db.Where("sm.menu_name like ?", "%"+selectSysMenu.MenuName+"%")
	}
	if selectSysMenu.Status != "" {
		db.Where("sm.status = ?", selectSysMenu.Status)
	}
	if selectSysMenu.UserId != 0 {
		db.Where("sur.user_id = ?", selectSysMenu.UserId)
	}
	if selectSysMenu.RoleId != 0 {
		db.Where("srm.role_id = ?", selectSysMenu.RoleId)
	}
	if !selectSysMenu.StartTime.IsZero() && !selectSysMenu.EndTime.IsZero() {
		db.Where("create_time BETWEEN ? AND ?", selectSysMenu.StartTime, selectSysMenu.EndTime)
	}
	db.Where("sm.deleted = 0").Order("sm.parent_id").Order("sm.order_num").Order("sm.create_time desc")
	err := db.Find(&sysMenuList).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetMenuById(menuId int) (sysMenu *model_sys.SysMenuVo) {
	db := mysql.Db().Model(&model_sys.SysMenuPo{})
	err := db.Where("menu_id = ? and deleted = 0", menuId).First(&sysMenu).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func SaveSysMenu(sysMenu *model_sys.SysMenuPo) (int64, int) {
	res := mysql.Db().Save(&sysMenu)
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected, sysMenu.MenuId
}

func IsChildren(menuId string) bool {
	var sysMenuList []*model_sys.SysMenuVo
	db := mysql.Db().Model(&model_sys.SysMenuPo{})
	err := db.Where("parent_id = ?", menuId).Find(&sysMenuList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return len(sysMenuList) > 0
}

// RemoveSysMenu 逻辑删除
func RemoveSysMenu(menuIds string, uk string) int64 {
	db := mysql.Db().Model(&model_sys.SysMenuPo{})
	db.Where("menu_id in (?)", menuIds)
	res := db.Updates(&model_sys.SysMenuPo{BaseEntity: model.BaseEntity{UpdateTime: utils.Time(time.Now()), UpdateBy: uk, Deleted: "1"}})
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

// DeletedSysMenu 删除
func DeletedSysMenu(menuIds string) int64 {
	db := mysql.Db().Model(&model_sys.SysMenuPo{})
	db.Where("menu_id in (?)", menuIds)
	res := db.Delete(&model_sys.SysMenuPo{})
	if res.Error != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", res.Error))
	}
	return res.RowsAffected
}

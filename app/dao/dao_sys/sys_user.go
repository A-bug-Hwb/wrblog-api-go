package dao_sys

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/mysql"
)

func GetUserByUk(uk string) (sysUser *model_sys.SysUserPo) {
	err := mysql.Db().Model(&model_sys.SysUserPo{}).Where("uk = ? and deleted = 0", uk).First(&sysUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetUserBySpaceUrl(spaceUrl string) (sysUser *model_sys.SysUserPo) {
	err := mysql.Db().Model(&model_sys.SysUserPo{}).Where("space_url = ? and deleted = 0", spaceUrl).First(&sysUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetUserByMk(mk string) (sysUser *model_sys.SysUserPo) {
	err := mysql.Db().Model(&model_sys.SysUserPo{}).Where("mobile = ? and deleted = 0", mk).First(&sysUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetUserByMbk(mbk string) (sysUser *model_sys.SysUserPo) {
	err := mysql.Db().Model(&model_sys.SysUserPo{}).Where("mailbox = ? and deleted = 0", mbk).First(&sysUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

func GetUserById(userId int) (sysUser *model_sys.SysUserVo) {
	err := mysql.Db().Model(&model_sys.SysUserPo{}).Where("user_id = ? and deleted = 0", userId).First(&sysUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}

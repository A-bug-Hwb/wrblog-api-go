package service_sys

import (
	"wrblog-api-go/app/dao/dao_sys"
	"wrblog-api-go/app/model/model_sys"
)

func GetUserByUk(uk string) *model_sys.SysUserPo {
	return dao_sys.GetUserByUk(uk)
}

func GetUserBySpaceUrl(spaceUrl string) (sysUser *model_sys.SysUserPo) {
	return dao_sys.GetUserBySpaceUrl(spaceUrl)
}

func GetUserByMk(mk string) (sysUser *model_sys.SysUserPo) {
	return dao_sys.GetUserByMk(mk)
}

func GetUserByMbk(mbk string) *model_sys.SysUserPo {
	return dao_sys.GetUserByUk(mbk)
}

func GetUserById(userId int) *model_sys.SysUserVo {
	return dao_sys.GetUserById(userId)
}

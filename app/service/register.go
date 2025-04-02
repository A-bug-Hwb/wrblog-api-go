package service

import (
	"errors"
	"wrblog-api-go/app/common/constants"
	"wrblog-api-go/app/common/token"
	"wrblog-api-go/app/common/utils"
	"wrblog-api-go/app/dao/dao_sys"
	"wrblog-api-go/pkg/redis"
)

func Register(registerForm *token.RegisterForm) (err error) {
	if err = registerForm.RegisterFormValidate(); err != nil {
		return
	}
	sysUser := dao_sys.GetUserByUk(registerForm.Uk)
	if sysUser != nil {
		err = errors.New("用户已存在！")
	} else if utils.VerifySpaceUrl(registerForm.Uk) {
		err = errors.New("不能用空间url的格式！")
	} else {
		ck, errCode := redis.Get(constants.CODE_KEY + registerForm.CkId)
		if errCode != nil {
			err = errors.New("验证码过期！")
		} else if string(ck) != registerForm.Ck {
			err = errors.New("验证码错误！")
		}
	}
	return
}

func RegisterKey(registerForm *token.RegisterForm) (err error) {
	if err = registerForm.RegisterFormValidate(); err != nil {
		return
	}
	privateKey, _ := redis.Get(constants.RSA_KEY + registerForm.RkId)
	pk := utils.RsaDecrypt(registerForm.Pk, string(privateKey))
	cpk := utils.RsaDecrypt(registerForm.Cpk, string(privateKey))
	registerForm.Pk = pk
	registerForm.Cpk = cpk
	sysUser := dao_sys.GetUserByUk(registerForm.Uk)
	if sysUser != nil {
		err = errors.New("用户已存在！")
	} else if utils.VerifySpaceUrl(registerForm.Uk) {
		err = errors.New("不能用空间url的格式！")
	} else {
		ck, errCode := redis.Get(constants.CODE_KEY + registerForm.CkId)
		if errCode != nil {
			err = errors.New("验证码过期！")
		} else if string(ck) != registerForm.Ck {
			err = errors.New("验证码错误！")
		}
	}
	return
}

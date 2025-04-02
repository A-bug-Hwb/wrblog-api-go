package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"wrblog-api-go/app/common/constants"
	"wrblog-api-go/app/common/token"
	utils2 "wrblog-api-go/app/common/utils"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/app/service/service_sys"
	"wrblog-api-go/config"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/redis"
	"wrblog-api-go/pkg/utils"
)

func Login(loginForm *token.LoginForm, clientInfo *token.ClientInfo) (tokenInfo *token.TokenInfo, err error) {
	val, _ := redis.Get(fmt.Sprintf("%s%s", constants.LOGIN_AK, loginForm.Ak))
	loginNum, _ := strconv.Atoi(string(val))
	if loginNum >= 5 {
		err = errors.New(fmt.Sprintf("密码错误次数超过五次，账户锁定%d分钟", constants.KEY_TIME))
		return
	}
	if err = loginForm.LoginFormValidate(); err != nil {
		return
	}
	var loginUser *token.LoginUser
	switch loginForm.Lt {
	case "1":
		loginUser, err = loginUk(loginForm, clientInfo)
		break
	case "2":
		loginUser, err = loginMk(loginForm, clientInfo)
		break
	case "3":
		loginUser, err = loginMbk(loginForm, clientInfo)
		break
	default:
		mylog.MyLog.Panic("不支持的登录方式")
		break
	}
	if err != nil {
		return
	}
	tokenInfo = token.CreateToken(loginUser)
	return
}

func LoginKey(loginForm *token.LoginForm, clientInfo *token.ClientInfo) (tokenInfo *token.TokenInfo, err error) {
	val, _ := redis.Get(fmt.Sprintf("%s%s", constants.LOGIN_AK, loginForm.Ak))
	loginNum, _ := strconv.Atoi(string(val))
	if loginNum >= 5 {
		err = errors.New(fmt.Sprintf("密码错误次数超过五次，账户锁定%d分钟", constants.KEY_TIME))
		return
	}
	if err = loginForm.LoginFormValidate(); err != nil {
		return
	}
	var loginUser *token.LoginUser
	switch loginForm.Lt {
	case "1":
		//账号密码登录逻辑
		privateKey, _ := redis.Get(constants.RSA_KEY + loginForm.RkId)
		redis.Del(constants.RSA_KEY + loginForm.RkId)
		loginForm.Pk = utils2.RsaDecrypt(loginForm.Pk, string(privateKey))
		loginUser, err = loginUk(loginForm, clientInfo)
		break
	case "2":
		loginUser, err = loginMk(loginForm, clientInfo)
		break
	case "3":
		loginUser, err = loginMbk(loginForm, clientInfo)
		break
	default:
		mylog.MyLog.Panic("不支持的登录方式")
	}
	if err != nil {
		return
	}
	tokenInfo = token.CreateToken(loginUser)
	return
}

// loginUk 账号密码登录
func loginUk(loginForm *token.LoginForm, clientInfo *token.ClientInfo) (loginUser *token.LoginUser, err error) {
	isVerify := service_sys.GetValueByKey("sys:account:captchaEnabled")
	if isVerify == "true" {
		data, _ := redis.Get(constants.CODE_KEY + loginForm.CkId)
		if len(data) == 0 {
			err = errors.New("验证码已过期！")
			return
		} else {
			if string(data) != loginForm.Ck {
				err = errors.New("验证码错误！")
			}
			return
		}
	}
	var sysUser *model_sys.SysUserPo
	sysUser = service_sys.GetUserByUk(loginForm.Ak)
	if sysUser == nil {
		if utils2.VerifySpaceUrl(loginForm.Ak) {
			sysUser = service_sys.GetUserBySpaceUrl(loginForm.Ak)
		} else if utils2.VerifyMobile(loginForm.Ak) {
			sysUser = service_sys.GetUserByMk(loginForm.Ak)
		} else if utils2.VerifyMailBox(loginForm.Ak) {
			sysUser = service_sys.GetUserByMbk(loginForm.Ak)
		}
		if sysUser == nil {
			err = errors.New("未查到账号信息！")
			return
		}
	}
	if !utils2.CheckPasswordHash(loginForm.Pk, sysUser.Pk) {
		val, _ := redis.Get(fmt.Sprintf("%s%s", constants.LOGIN_AK, loginForm.Ak))
		loginNum, _ := strconv.Atoi(string(val))
		redis.SetTime(fmt.Sprintf("%s%s", constants.LOGIN_AK, loginForm.Ak), loginNum+1, constants.KEY_TIME*time.Minute)
		err = errors.New("账号密码错误！")
		return
	}
	loginUser = getLoginUser(loginForm, clientInfo, sysUser)
	return
}

// LoginMk 手机号验证码登录
func loginMk(loginForm *token.LoginForm, clientInfo *token.ClientInfo) (loginUser *token.LoginUser, err error) {
	ck, _ := redis.Get(constants.CODE_KEY + loginForm.CkId)
	redis.Del(constants.CODE_KEY + loginForm.CkId)
	if len(ck) == 0 {
		err = errors.New("验证码已过期！")
	} else if string(ck) != loginForm.Ck {
		err = errors.New("验证码错误！")
	} else {
		var sysUser *model_sys.SysUserPo
		sysUser = service_sys.GetUserByMk(loginForm.Ak)
		if sysUser == nil {
			err = errors.New("未查到账号信息！")
		} else {
			loginUser = getLoginUser(loginForm, clientInfo, sysUser)
		}
	}
	return
}

// LoginMbk 邮箱验证码登录
func loginMbk(loginForm *token.LoginForm, clientInfo *token.ClientInfo) (loginUser *token.LoginUser, err error) {
	ck, err := redis.Get(constants.CODE_KEY + loginForm.CkId)
	if err != nil {
		return
	}
	redis.Del(constants.CODE_KEY + loginForm.CkId)
	if len(ck) == 0 {
		err = errors.New("验证码已过期！")
	} else if string(ck) != loginForm.Ck {
		err = errors.New("验证码错误！")
	} else {
		var sysUser *model_sys.SysUserPo
		sysUser = service_sys.GetUserByMbk(loginForm.Ak)
		if sysUser == nil {
			err = errors.New("未查到账号信息！")
		} else {
			loginUser = getLoginUser(loginForm, clientInfo, sysUser)
		}
	}
	return
}

func getLoginUser(loginForm *token.LoginForm, clientInfo *token.ClientInfo, sysUser *model_sys.SysUserPo) *token.LoginUser {
	return &token.LoginUser{
		UserId:     sysUser.UserId,
		Uk:         sysUser.Uk,
		Ak:         loginForm.Ak,
		LoginTime:  utils.GetNowDate(),
		ExpireTime: config.Conf.ConfigInfo.Token.ExpireTime,
		ClientInfo: *clientInfo,
	}
}

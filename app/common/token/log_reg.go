package token

import (
	"encoding/json"
	"errors"
	"strings"
	"wrblog-api-go/app/common/utils"
)

type LoginForm struct {
	Ak   string `json:"ak"`   //账号（必填）
	Lt   string `json:"lt"`   //登录类型（必填）1账号密码，2短信验证码，3邮箱验证码，4微信扫码
	RkId string `json:"rkId"` //rsa秘钥id（加密传输的秘钥id索引）
	Pk   string `json:"pk"`   //密码（账号密码登录必填）
	CkId string `json:"ckId"` //验证码id（验证码对应的id索引）
	Ck   string `json:"ck"`   //验证码（登录时填写图片验证码，短信和邮箱则填写发送的验证码）
}

func (loginForm *LoginForm) LoginFormValidate() (err error) {
	switch loginForm.Lt {
	case "1":
		if strings.TrimSpace(loginForm.Ak) == "" {
			err = errors.New("账号不能为空！")
		} else if strings.TrimSpace(loginForm.Pk) == "" {
			err = errors.New("密码不能为空！")
		}
		break
	case "2":
		if strings.TrimSpace(loginForm.Ak) == "" {
			err = errors.New("手机号不能为空！")
		} else if !utils.VerifyMobile(loginForm.Ak) {
			err = errors.New("手机号格式不正确！")
		} else if strings.TrimSpace(loginForm.Ck) == "" {
			err = errors.New("验证码不能为空！")
		}
		break
	case "3":
		if strings.TrimSpace(loginForm.Ak) == "" {
			err = errors.New("邮箱不能为空！")
		} else if !utils.VerifyMailBox(loginForm.Ak) {
			err = errors.New("邮箱格式不正确！")
		} else if strings.TrimSpace(loginForm.Ck) == "" {
			err = errors.New("验证码不能为空！")
		}
		break
	default:
		err = errors.New("不支持该登录方式！")
	}
	return
}

type LoginUser struct {
	UserId      int      `json:"userId"`
	UserKey     string   `json:"userKey"`
	Ak          string   `json:"ak"`
	Uk          string   `json:"uk"`
	LoginTime   string   `json:"loginTime"`
	ExpireTime  int      `json:"expireTime"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
	ClientInfo
}

// 序列化
func (loginUser *LoginUser) MarshalBinary() (data []byte, err error) {
	return json.Marshal(loginUser)
}

// 反序列化
func (loginUser *LoginUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, loginUser)
}

type RegisterForm struct {
	Uk     string `json:"uk"`
	Pk     string `json:"pk"`
	Cpk    string `json:"cpk"`
	RkId   string `json:"rkId"`
	Mobile string `json:"mobile"`
	CkId   string `json:"ckId"`
	Ck     string `json:"ck"`
}

func (registerForm *RegisterForm) RegisterFormValidate() (err error) {
	if strings.TrimSpace(registerForm.Uk) == "" {
		err = errors.New("账号不能为空！")
	} else if strings.TrimSpace(registerForm.Pk) == "" {
		err = errors.New("密码不能为空！")
	} else if strings.TrimSpace(registerForm.Cpk) == "" {
		err = errors.New("确认密码不能为空！")
	} else if strings.TrimSpace(registerForm.Pk) != strings.TrimSpace(registerForm.Cpk) {
		err = errors.New("两次密码不一致！")
	} else if strings.TrimSpace(registerForm.Mobile) == "" {
		err = errors.New("手机号不能为空！")
	} else if strings.TrimSpace(registerForm.Ck) == "" {
		err = errors.New("验证码不能为空！")
	}
	return
}

type TokenInfo struct {
	Token      string `json:"token"`
	ResToken   string `json:"resToken"`
	ExpireTime int    `json:"expireTime"`
}

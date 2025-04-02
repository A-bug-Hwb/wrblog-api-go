package service

import (
	"fmt"
	"time"
	"wrblog-api-go/app/common/constants"
	"wrblog-api-go/app/common/utils"
	"wrblog-api-go/app/dao/dao_sys"
	"wrblog-api-go/config"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/redis"
)

func GetImgCode() (data map[string]any) {
	data = make(map[string]any)
	isVerify := dao_sys.GetValueByKey("sys:account:captchaEnabled")
	if isVerify == "true" {
		id, b64s, answer, errCode := utils.CreateCode()
		if errCode != nil {
			mylog.MyLog.Panic(fmt.Sprintf("获取登录验证码，验证码生成失败:%s", errCode.Error()))
		}
		redis.SetTime(constants.CODE_KEY+id, answer, constants.KEY_TIME*time.Minute)
		data["ckId"] = id
		data["image"] = b64s
		data["isVerify"] = true
	} else {
		data["isVerify"] = false
	}
	return
}

func GetPublicKey() (data map[string]string) {
	isVerify := dao_sys.GetValueByKey("sys:key:generateKey")
	data = make(map[string]string)
	rkId := utils.GetUUIDString()
	data["rkId"] = rkId
	publicKey, privateKey := config.Conf.ConfigInfo.Key.PublicKey, config.Conf.ConfigInfo.Key.PrivateKey
	if isVerify == "true" {
		publicKey, privateKey = utils.RsaGenKey()
	}
	data["publicKey"] = publicKey
	redis.SetTime(constants.RSA_KEY+rkId, privateKey, constants.KEY_TIME*time.Minute)
	return
}

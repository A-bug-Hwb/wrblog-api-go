package token

import (
	"fmt"
	"strconv"
	"time"
	"wrblog-api-go/app/common/constants"
	"wrblog-api-go/app/common/utils"
	"wrblog-api-go/config"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/redis"
)

var authJwt = NewJWT()

func CreateToken(loginUser *LoginUser) (tokenInfo *TokenInfo) {
	var userKey = utils.GetUUIDString()
	loginUser.UserKey = userKey
	claims := &CustomClaims{
		Uid:    userKey,
		UserId: strconv.Itoa(loginUser.UserId),
	}
	//将用户信息存入缓存
	RefreshToken(userKey, loginUser)
	token, err := authJwt.GenerateToken(claims)
	if err != nil {
		mylog.MyLog.Panic("Token生成失败！")
	}
	tokenInfo = &TokenInfo{
		Token:      token,
		ExpireTime: config.Conf.ConfigInfo.Token.ExpireTime,
	}
	return
}

// GetTokenVal 获取token信息
func GetTokenVal(token string) (claims *CustomClaims, ok bool) {
	return authJwt.ParseToken(token)
}

// RefreshToken 刷新token
func RefreshToken(userKey any, loginUser *LoginUser) {
	redis.SetTime(fmt.Sprintf("%s%s:%s", constants.LOGIN_USER_KEY, strconv.Itoa(loginUser.UserId), userKey), loginUser, time.Duration(config.Conf.ConfigInfo.Token.ExpireTime)*time.Minute)
}

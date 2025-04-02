package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"wrblog-api-go/app/common/constants"
	"wrblog-api-go/config"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/redis"
)

type ClientInfo struct {
	IpAddr        string `json:"ipAddr" form:"ipAddr"`
	LoginLocation string `json:"loginLocation" form:"loginLocation"`
	Browser       string `json:"browser" form:"browser"`
	Os            string `json:"os" form:"os"`
}

// GetToken 获取请求体中的token
func GetToken(c *gin.Context) string {
	tokenInfo := c.GetHeader(config.Conf.ConfigInfo.Token.Header)
	if tokenInfo != "" && strings.Contains(tokenInfo, constants.TOKEN_PREFIX) {
		return tokenInfo[len(constants.TOKEN_PREFIX):]
	} else {
		return tokenInfo
	}
}

func GetLoginUser(c *gin.Context) (loginUser *LoginUser, err error) {
	claims, ok := GetTokenVal(GetToken(c))
	if !ok {
		err = errors.New("登录过期")
		return
	}
	info, err := redis.Get(fmt.Sprintf("%s%s:%s", constants.LOGIN_USER_KEY, claims.UserId, claims.Uid))
	if err != nil {
		err = errors.New("登录过期")
		return
	}
	err = json.Unmarshal(info, &loginUser)
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("Json转换失败：%s", err))
	}
	RefreshToken(claims.Uid, loginUser)
	return
}

func GetUserId(c *gin.Context) int {
	loginUser, err := GetLoginUser(c)
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("获取登录UserId失败：%s", err))
	}
	return loginUser.UserId
}

func IsAdmin(userId int) bool {
	return userId == 888888888888888888
}

func GetUk(c *gin.Context) string {
	loginUser, err := GetLoginUser(c)
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("获取登录Uk失败：%s", err))
	}
	return loginUser.Uk
}

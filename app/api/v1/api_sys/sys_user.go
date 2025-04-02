package api_sys

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wrblog-api-go/app/common/token"
	"wrblog-api-go/app/dao/dao_sys"
	"wrblog-api-go/pkg/result"
)

// @Tags  System - 用户管理
// @Summary  获取当前用户的登录信息
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} result.Result "OK"
// @Router /sys/getLoginUser [get]
func ApiLoginUser(c *gin.Context) {
	loginUser, _ := token.GetLoginUser(c)
	c.JSON(http.StatusOK, result.Ok(loginUser))
}

// @Tags  System - 用户管理
// @Summary  获取当前登录用户信息
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} result.Result "OK"
// @Router /sys/getUserInfo [get]
func ApiUserInfo(c *gin.Context) {
	data := make(map[string]any)
	sysUser := dao_sys.GetUserById(token.GetUserId(c))
	data["user"] = sysUser
	c.JSON(http.StatusOK, result.Ok(data))
}

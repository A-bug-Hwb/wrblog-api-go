package api_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wrblog-api-go/app/common/token"
	"wrblog-api-go/app/service"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/result"
)

// @Tags  Auth - 认证授权
// @Summary  登录接口（无需加密）
// @Accept json
// @Produce json
// @Param registerForm body token.RegisterForm true "registerForm"
// @Success 200 {object} result.Result "OK"
// @Router /auth/register [post]
func ApiRegister(c *gin.Context) {
	var registerForm *token.RegisterForm
	if err := c.ShouldBindBodyWithJSON(&registerForm); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err.Error()))
	}
	var res *result.Result
	err := service.Register(registerForm)
	if err != nil {
		res = result.Fail(fmt.Sprintf("注册失败，%s", err.Error()))
	} else {
		res = result.Ok(nil)
	}
	c.JSON(http.StatusOK, res)
}

// @Tags  Auth - 认证授权
// @Summary  登录接口（加密）
// @Accept json
// @Produce json
// @Param registerForm body token.RegisterForm true "registerForm"
// @Success 200 {object} result.Result "OK"
// @Router /auth/registerKey [post]
func ApiRegisterKey(c *gin.Context) {
	var registerForm *token.RegisterForm
	if err := c.ShouldBindBodyWithJSON(&registerForm); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err.Error()))
	}
	var res *result.Result
	err := service.RegisterKey(registerForm)
	if err != nil {
		res = result.Fail(fmt.Sprintf("注册失败，%s", err.Error()))
	} else {
		res = result.Ok(nil)
	}
	c.JSON(http.StatusOK, res)
}

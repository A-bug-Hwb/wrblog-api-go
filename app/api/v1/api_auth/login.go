package api_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"wrblog-api-go/app/common/constants"
	"wrblog-api-go/app/common/token"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/app/service"
	"wrblog-api-go/app/service/service_sys"
	"wrblog-api-go/pkg/client"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/redis"
	"wrblog-api-go/pkg/result"
	"wrblog-api-go/pkg/utils"
)

// @Tags  Auth - 认证授权
// @Summary  登录接口（无需加密）
// @Accept json
// @Produce json
// @Param loginForm body token.LoginForm true "loginForm"
// @Success 200 {object} result.Result "OK"
// @Router /auth/login [post]
func ApiLogin(c *gin.Context) {
	var loginForm *token.LoginForm
	if err := c.ShouldBindBodyWithJSON(&loginForm); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err.Error()))
	}
	clientInfo := client.GetClient(c)
	tokenInfo, err := service.Login(loginForm, clientInfo)
	logininfor := &model_sys.SysLogininfor{
		Ak:         loginForm.Ak,
		ClientInfo: *clientInfo,
		AccessTime: utils.Time(time.Now()),
	}
	var res *result.Result
	if err != nil {
		res = result.Fail(fmt.Sprintf("登录失败，%s", err))
		logininfor.Status = "1"
		logininfor.Msg = fmt.Sprintf("登录失败，%s", err)
	} else {
		res = result.Ok(tokenInfo)
		logininfor.Status = "0"
		logininfor.Msg = "登录成功！"
	}
	go service_sys.AddLoginifor(logininfor)
	c.JSON(http.StatusOK, res)
}

// @Tags  Auth - 认证授权
// @Summary  登录接口（加密）
// @Accept json
// @Produce json
// @Param loginForm body token.LoginForm true "loginForm"
// @Success 200 {object} result.Result "OK"
// @Router /auth/loginKey [post]
func ApiLoginKey(c *gin.Context) {
	var loginForm *token.LoginForm
	if err := c.ShouldBindBodyWithJSON(&loginForm); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err.Error()))
	}
	val, _ := redis.Get(fmt.Sprintf("%s%s", constants.LOGIN_AK, loginForm.Ak))
	loginNum, _ := strconv.Atoi(string(val))
	if loginNum >= 5 {
		res := result.Fail(fmt.Sprintf("密码错误次数超过五次，账户锁定%d分钟", constants.KEY_TIME))
		c.JSON(http.StatusOK, res)
		return
	}
	clientInfo := client.GetClient(c)
	tokenInfo, err := service.LoginKey(loginForm, clientInfo)
	logininfor := &model_sys.SysLogininfor{
		Ak:         loginForm.Ak,
		ClientInfo: *clientInfo,
		AccessTime: utils.Time(time.Now()),
	}
	var res *result.Result
	if err != nil {
		res = result.Fail(fmt.Sprintf("登录失败，%s", err))
		logininfor.Status = "1"
		logininfor.Msg = fmt.Sprintf("登录失败，%s", err.Error())
		redis.SetTime(fmt.Sprintf("%s%s", constants.LOGIN_AK, loginForm.Ak), loginNum+1, constants.KEY_TIME*time.Minute)
	} else {
		res = result.Ok(tokenInfo)
		logininfor.Status = "0"
		logininfor.Msg = "登录成功！"
	}
	go service_sys.AddLoginifor(logininfor)
	c.JSON(http.StatusOK, res)
}

// @Tags  Auth - 认证授权
// @Summary  退出登录
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} result.Result "OK"
// @Router /auth/logout [get]
func ApiLogout(c *gin.Context) {
	loginUser, _ := token.GetLoginUser(c)
	if loginUser != nil {
		redis.Del(fmt.Sprintf("%s%s:%s", constants.LOGIN_USER_KEY, strconv.Itoa(loginUser.UserId), loginUser.UserKey))
	}
	c.JSON(http.StatusOK, result.Ok(nil))
}

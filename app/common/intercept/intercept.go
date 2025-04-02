package intercept

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wrblog-api-go/app/common/token"
	"wrblog-api-go/pkg/constants"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/result"
)

// CheckHttp 拦截请求、并构建可重复读body
func CheckHttp(c *gin.Context) {
	mylog.MyLog.Debug(fmt.Sprintf("有请求发送------>%s------>%s", c.Request.Method, c.Request.RequestURI))
	if (c.Request.Method == "POST" || c.Request.Method == "PUT") && c.ContentType() == "application/json" {
		var body = make(map[string]any)
		if err := c.ShouldBindBodyWithJSON(&body); err != nil {
			mylog.MyLog.Debug(fmt.Sprintf("未读取到Body参数------>%s--->%s", body, err.Error()))
		}
		mylog.MyLog.Debug(fmt.Sprintf("Body参数------>%s", body))
	} else if c.Request.Method == "GET" || c.Request.Method == "DELETE" {
		var params = make(map[string]string)
		if err := c.ShouldBindQuery(&params); err != nil {
			mylog.MyLog.Debug(fmt.Sprintf("未读取到Params参数------>%s--->%s", params, err.Error()))
		}
		mylog.MyLog.Debug(fmt.Sprintf("Params参数------>%s", params))
	}
	c.Next()
	//全局响应，如果状态不对，则报系统错误
	status := c.Writer.Status()
	if status != http.StatusOK {
		c.JSON(http.StatusOK, result.Fail("系统错误，请联系管理员"))
	}
}

// CheckIp ip拦截
func CheckIp(c *gin.Context) {
	//ip白名单
	if NotIp(c.ClientIP()) {
		c.Next()
	} else {
		mylog.MyLog.Panic("该Ip无权访问:%s", c.ClientIP())
		c.JSON(http.StatusOK, result.New(constants.IP_LOCKED, fmt.Sprintf("该Ip无权访问:%s", c.ClientIP()), nil))
		c.Abort() //停止执行
	}
}

// CheckToken token拦截
func CheckToken(c *gin.Context) {
	path := c.Request.RequestURI
	//判断拦截的路由
	if !NotIntercept(path) {
		loginUser, err := token.GetLoginUser(c)
		if err != nil {
			res := result.New(constants.UNAUTHORIZED, fmt.Sprintf("请求访问：%s,认证失败,%s！", path, err.Error()), nil)
			c.JSON(http.StatusOK, res)
			c.Abort() //停止执行
			return
		}
		ctx := context.WithValue(context.Background(), "loginUser", loginUser)
		c.Set("ctx", ctx)
	}
	c.Next() //继续执行
}

// CrossOriginMiddleware 跨域设置
func CrossOriginMiddleware(c *gin.Context) {
	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}

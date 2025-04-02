package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"wrblog-api-go/app/common/intercept"
	"wrblog-api-go/app/router/v1"
	"wrblog-api-go/config"
	router2 "wrblog-api-go/demo/router"
	"wrblog-api-go/pkg/constants"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/result"
	"wrblog-api-go/pkg/swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//注册swagger
	swagger.RegisterSwagger(router)
	//处理找不到路由
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	//处理发生异常
	router.Use(Recover)
	//跨域配置
	router.Use(intercept.CrossOriginMiddleware)
	//全局拦截
	router.Use(intercept.CheckHttp, intercept.CheckIp, intercept.CheckToken)
	//ico配置
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "欢迎访问golang gin框架")
	})
	//接口
	apiV1 := router.Group("/api/v1")
	//静态资源映射
	apiV1.Static(config.Conf.ConfigInfo.FilePrefix, config.Conf.ConfigInfo.Profile)
	v1.WrAuthApi(apiV1)
	v1.WrSystemApi(apiV1)
	v1.WrBaseApi(apiV1)
	v1.WrFileApi(apiV1)
	//demo
	router2.ApiDemo(apiV1)
	return router
}

func HandleNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, result.New(constants.NOT_FOUND, "系统接口404！", nil))
}

// Recover 内部发生异常时的处理
func Recover(c *gin.Context) {
	defer func() {
		if info := recover(); info != nil {
			val := reflect.ValueOf(info)
			switch val.Type().String() {
			case "string":
				mylog.MyLog.Error(fmt.Sprintf("服务器内部错误！%s", val))
				c.JSON(http.StatusOK, result.Fail(fmt.Sprintf("服务器内部错误！%s", val)))
				panic(val)
			case "*logrus.Entry":
				log := info.(*logrus.Entry)
				mylog.MyLog.Error(fmt.Sprintf("服务器内部错误！%s", log.Message))
				c.JSON(http.StatusOK, result.Fail(fmt.Sprintf("服务器内部错误！%s", log.Message)))
				panic(log.Message)
			default:
				mylog.MyLog.Error(fmt.Sprintf("服务器内部错误！%s", info))
				c.JSON(http.StatusOK, result.Fail(fmt.Sprintf("服务器内部错误！%s", info)))
				panic(info)
			}
		}
	}()
	//继续后续接口调用
	c.Next()
}

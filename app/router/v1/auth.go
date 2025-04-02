package v1

import (
	"github.com/gin-gonic/gin"
	"wrblog-api-go/app/api/v1/api_auth"
)

func WrAuthApi(apiV1 *gin.RouterGroup) {
	//获取列表
	AuthApi := apiV1.Group("/auth")
	{
		AuthApi.GET("/getImgCode", api_auth.ApiImgCode)
		AuthApi.GET("/getPublicKey", api_auth.ApiPublicKey)
		AuthApi.POST("/login", api_auth.ApiLogin)
		AuthApi.POST("/loginKey", api_auth.ApiLoginKey)
		AuthApi.GET("/logout", api_auth.ApiLogout)
		AuthApi.POST("/register", api_auth.ApiRegister)
		AuthApi.POST("/registerKey", api_auth.ApiRegisterKey)
	}
}

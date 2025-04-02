package v1

import (
	"github.com/gin-gonic/gin"
	"wrblog-api-go/app/api/v1/api_sys"
)

func WrSystemApi(apiV1 *gin.RouterGroup) {
	//获取列表
	SysApi := apiV1.Group("/sys")
	{
		SysApi.GET("/getLoginUser", api_sys.ApiLoginUser)
		SysApi.GET("/getUserInfo", api_sys.ApiUserInfo)
		SysApi.GET("/getRouters", api_sys.ApiRouters)
		SysMenuApi := SysApi.Group("/sysMenu")
		{
			SysMenuApi.GET("/getTree", api_sys.ApiSysMenuTree)
			SysMenuApi.GET("/:menuId", api_sys.ApiSysMenuById)
			SysMenuApi.POST("/add", api_sys.ApiAddSysMenu)
			SysMenuApi.PUT("/edit", api_sys.ApiEditSysMenu)
			SysMenuApi.DELETE("/remove", api_sys.ApiRemoveSysMenu)
			SysMenuApi.DELETE("/delete", api_sys.ApiDeleteSysMenu)
		}
		SysDictApi := SysApi.Group("/sysDict")
		{
			TypeApi := SysDictApi.Group("/type")
			{
				TypeApi.GET("/page", api_sys.ApiDictTypePageList)
				TypeApi.GET("/:dictId", api_sys.ApiDictTypeById)
				TypeApi.POST("/add", api_sys.ApiAddDictType)
				TypeApi.PUT("/edit", api_sys.ApiEditDictType)
				TypeApi.DELETE("/removes", api_sys.ApiRemoveDictType)
				TypeApi.DELETE("/refreshCache", api_sys.ApiRemoveCacheAll)
			}
			DataApi := SysDictApi.Group("/data")
			{
				DataApi.GET("/page", api_sys.ApiDictDataPageList)
				DataApi.GET("/list/:dictType", api_sys.ApiDictDataByType)
				DataApi.GET("/:dictCode", api_sys.ApiDictDataByCode)
				DataApi.POST("/add", api_sys.ApiAddDictData)
				DataApi.PUT("/edit", api_sys.ApiEditDictData)
				DataApi.DELETE("/removes", api_sys.ApiRemoveDictData)
			}
		}
	}
}

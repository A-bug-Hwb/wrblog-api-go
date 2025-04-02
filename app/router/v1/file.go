package v1

import (
	"github.com/gin-gonic/gin"
	"wrblog-api-go/app/api/v1/api_sys"
)

func WrFileApi(apiV1 *gin.RouterGroup) {
	//获取列表
	FileApi := apiV1.Group("/file")
	{
		FileApi.POST("/upload", api_sys.ApiUpload)
	}
}

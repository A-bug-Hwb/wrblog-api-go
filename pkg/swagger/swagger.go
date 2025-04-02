package swagger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strings"
	"wrblog-api-go/config"
	"wrblog-api-go/docs"
)

func RegisterSwagger(r gin.IRouter) {
	// API文档访问地址: http://host/swagger/index.html
	// 注解定义可参考 https://github.com/swaggo/swag#declarative-comments-format
	// 样例 https://github.com/swaggo/swag/blob/master/example/basic/api/api.go
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "管理后台接口"
	docs.SwaggerInfo.Description = "实现一个管理系统的后端API服务"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", config.Conf.Server.Host, config.Conf.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//转为小写
	mode := strings.ToLower(config.Conf.Server.Mode)
	if mode == "debug" || mode == "test" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	} else {
		r.GET("/swagger/*any", func(c *gin.Context) {
			c.JSON(http.StatusOK, "当前环境未开启Swagger功能！")
		})
	}
}

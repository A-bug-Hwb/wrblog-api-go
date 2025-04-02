package router

import (
	"github.com/gin-gonic/gin"
	"wrblog-api-go/demo/api/demo/api_rabbit"
	"wrblog-api-go/demo/modules/socket"
)

func ApiDemo(apiV1 *gin.RouterGroup) {
	DemoApi := apiV1.Group("/demo")
	{
		SocketApi := DemoApi.Group("/socket")
		{
			SocketApi.GET("/open", socket.WebsocketHandler)
		}
		RocketApi := DemoApi.Group("/rabbit")
		{
			RocketApi.GET("/start", api_rabbit.Start)
			RocketApi.POST("/push", api_rabbit.PushData)
		}
	}

}

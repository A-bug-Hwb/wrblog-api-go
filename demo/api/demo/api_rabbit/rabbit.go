package api_rabbit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wrblog-api-go/demo/modules/rabbit"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/result"
)

// @Tags  Demo - Rabbit
// @Summary  Rabbit开始执行
// @Accept json
// @Produce json
// @Success 200 {object} result.Result "OK"
// @Router /demo/rabbit/start [get]
func Start(c *gin.Context) {
	go rabbit.ReceiveAll()
	c.JSON(http.StatusOK, result.Ok("开启"))
}

// @Tags  Demo - Rabbit
// @Summary  Rabbit发送数据
// @Accept json
// @Produce json
// @Param rabbitData body rabbit.RabbitData true "rabbitData"
// @Success 200 {object} result.Result "OK"
// @Router /demo/rabbit/push [post]
func PushData(c *gin.Context) {
	var rbData *rabbit.RabbitData
	if err := c.ShouldBindBodyWithJSON(&rbData); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err.Error()))
	}
	var res *result.Result
	var err error
	switch rbData.Type {
	case rabbit.None:
		err = rabbit.NoneExchange(rbData)
		break
	case rabbit.Direct:
		err = rabbit.DirectExchange(rbData)
		break
	case rabbit.Topic:
		err = rabbit.NoneExchange(rbData)
		break
	case rabbit.Fanout:
		err = rabbit.NoneExchange(rbData)
		break
	case rabbit.Headers:
		err = rabbit.NoneExchange(rbData)
		break
	}
	if err != nil {
		res = result.Fail(fmt.Sprintf("rabbit消息发送失败：模式：%s，信息：%s", rbData.Type, err))
	} else {
		res = result.Ok("发送成功")
	}
	c.JSON(http.StatusOK, res)
}

package socket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"wrblog-api-go/app/common/utils"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/redis"
)

var wsKey = "socket_key:" //通信socket_key,存储客户端在线状态
var timeout = 2           //超时时间，{timeout}分钟没有心跳则下线

var clientMap = make(map[string]*client)

type client struct {
	UserId   string          `json:"userId"`
	ClientId string          `json:"clientId"`
	OS       string          `json:"oS"`
	Ip       string          `json:"ip"`
	Ws       *websocket.Conn `json:"ws"`
}

type socketData struct {
	UserId  string `json:"userId"`
	Type    string `json:"type"`
	Content struct {
		UserId string `json:"userId"`
		Type   string `json:"type"`
		Msg    string `json:"msg"`
	} `json:"content"`
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 这里可以根据实际情况修改，例如只允许特定来源的请求升级为WebSocket连接
	},
}

func WebsocketHandler(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		err = receive(ws)
		if err != nil {
			break
		}
	}
}

// 接收消息处理
func receive(ws *websocket.Conn) (err error) {
	//读取ws中的数据
	mt, content, err := ws.ReadMessage()
	//断开连接
	if err != nil {
		return err
	}
	var data *socketData
	err = json.Unmarshal(content, &data)
	if err != nil {
		mylog.MyLog.Panic("消息体{%s}格式错误，拒绝处理！%s", string(content), err)
		return err
	}
	if data.Type == "heartbeat" {
		//来源客户端键值
		srcKey := fmt.Sprintf("%s%s", wsKey, data.UserId)
		clientInfo := &client{
			ClientId: utils.GetUUIDString(),
			UserId:   data.UserId,
			OS:       "pc",
			Ws:       ws,
		}
		//用户上线，存储在线客户端信息，并设置两分钟没收到心跳自动下线
		redis.SetTime(srcKey, clientInfo.ClientId, time.Duration(timeout)*time.Minute)
		//存储客户端信息
		clientMap[srcKey] = clientInfo
		return nil
	} else {
		//目标客户端键值
		srcKey := fmt.Sprintf("%s%s", wsKey, data.UserId)
		targetKey := fmt.Sprintf("%s%s", wsKey, data.Content.UserId)
		targetClient := clientMap[targetKey]
		if targetClient == nil {
			srcClient := clientMap[srcKey]
			data.Content.Msg = "目标用户不在线！"
			dataByte, _ := json.Marshal(data)
			err = push(mt, dataByte, srcClient.Ws)
		} else {
			err = push(mt, content, targetClient.Ws)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// 推送消息处理
func push(mt int, content []byte, ws *websocket.Conn) error {
	//写入ws数据
	err := ws.WriteMessage(mt, content)
	return err
}

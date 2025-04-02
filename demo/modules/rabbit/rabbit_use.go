package rabbit

import (
	"fmt"
	"wrblog-api-go/pkg/mylog"
)

func ReceiveAll() {
	ReceiveNoneExchange()
}

// 不使用交换机
func ReceiveNoneExchange() {
	//创建一个管道
	rabbit := NewRabbit()
	defer rabbit.Close()
	ch, err := rabbit.Channel()
	if err != nil {
		mylog.MyLog.Panic("rabbit消费者管道错误：模式：%s，信息：%s", None, err)
	}
	defer ch.Close()
	//为消息队列注册消费者
	//接收消息
	forever := make(chan bool)
	msgs, err := ch.Consume(
		"default", //消息队列的名称
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		mylog.MyLog.Panic("rabbit消费者错误：模式：%s，信息：%s", None, err)
	}
	go func() {
		for msg := range msgs {
			fmt.Println("Received a message: %s", msg.Body)
		}
	}()
	<-forever
}

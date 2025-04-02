package rabbit

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

// 不使用交换机
func NoneExchange(data *RabbitData) error {
	//创建一个管道
	rabbit := NewRabbit()
	defer rabbit.Close()
	ch, err := rabbit.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	//声明队列，如果队列不存在则自动创建，存在则跳过创建
	q, err := ch.QueueDeclare(
		"default", //消息队列的名称
		false,     //是否持久化
		false,     //是否自动删除
		false,     //是否具有排他性(仅创建它的程序才可用)
		false,     //是否阻塞处理
		nil,       //额外的属性
	)
	if err != nil {
		return err
	}
	body, _ := json.Marshal(data)
	err = ch.Publish(
		"",
		q.Name,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	return err
}

//1. Direct Exchange：
//直接交换是最常用的交换类型。消息从发送方到队列的路由过程完全依赖于路由键。如果一个队列绑定到某个路由键上，那么所有发送到这个路由键的消息都会被转发到这个队列

func DirectExchange(data *RabbitData) error {
	////创建一个管道
	//rabbit := NewRabbit()
	////定义队列的名称
	//queueNames := []string{"direct_Queue1", "direct_Queue2", "direct_Queue3", "direct_Queue4"}
	return nil
}

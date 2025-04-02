package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
	"gopkg.in/yaml.v3"
	"os"
	"wrblog-api-go/config"
)

// 不使用交换机
const None = "none"

// 四种交换机
const Direct = "direct"
const Topic = "topic"
const Fanout = "fanout"
const Headers = "headers"

// 消息体格式
type RabbitData struct {
	UserId  string `json:"userId"`
	Type    string `json:"type"`
	Content struct {
		UserId string `json:"userId"`
		Type   string `json:"type"`
		Msg    string `json:"msg"`
	} `json:"content"`
}

// RabbitConf 日志映射
type rabbitConf struct {
	Rabbit struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		VHost    string `yaml:"v-host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"rabbit"`
}

func NewRabbit() *amqp.Connection {
	//解析日志配置文件
	rbConf := analysisRabbit()
	//初始化日志配置信息
	rabbit, err := rbConf.initRabbit()
	if err != nil {
		//mylog.MyLog.Panic("Rabbit Connection initialization failed:%s", err.Error())
		panic(fmt.Sprintf("Rabbit Connection initialization failed:%s", err.Error()))
	} else {
		fmt.Println("Rabbit configuration initialization successful......")
	}
	return rabbit
}

func analysisRabbit() *rabbitConf {
	var rbConf *rabbitConf
	// 打开 YAML 文件
	file, err := os.Open(config.ConfPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse configuration file:%s", err.Error()))
	}
	defer file.Close()
	// 创建解析器
	decoder := yaml.NewDecoder(file)
	// 解析 YAML 数据
	err = decoder.Decode(&rbConf)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse log configuration file:%s", err.Error()))
	}
	return rbConf
}

func (rbConf *rabbitConf) initRabbit() (rabbit *amqp.Connection, err error) {
	rabbit, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d%s",
		rbConf.Rabbit.Username,
		rbConf.Rabbit.Password,
		rbConf.Rabbit.Host,
		rbConf.Rabbit.Port,
		rbConf.Rabbit.VHost))
	if err != nil {
		return nil, err
	}
	//开启消费
	//ReceiveAll()
	return rabbit, err
}

package mylog

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"strings"
	"time"
	"wrblog-api-go/config"
)

// LoggerConf 日志映射
type loggerConf struct {
	Logger struct {
		LogPath      string `yaml:"log-path"`
		Level        string `yaml:"level"`
		MaxAge       int    `yaml:"max-age"`
		RotationTime int    `yaml:"rotation-time"`
	} `yaml:"logger"`
}

var MyLog *logrus.Entry

var Logger = logrus.New()

func init() {
	//解析日志配置文件
	logConf := analysisLogger()
	//初始化日志配置信息
	logConf.initLogger()
	fmt.Println("Log configuration initialization successful......")
}

func analysisLogger() (logConf *loggerConf) {
	// 打开 YAML 文件
	file, err := os.ReadFile(config.ConfPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse configuration file:%s", err.Error()))
	}
	err = yaml.Unmarshal(file, &logConf)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse log configuration file:%s", err.Error()))
	}
	return
}

func (logConf *loggerConf) initLogger() {
	formatter := &logrus.TextFormatter{
		ForceColors: true, // 强制启用彩色输出
	}
	Logger.SetFormatter(formatter)
	//设置最低日志级别
	level := strings.ToLower(logConf.Logger.Level)
	switch level {
	case "panic":
		Logger.SetLevel(logrus.PanicLevel)
		break
	case "fatal":
		Logger.SetLevel(logrus.FatalLevel)
		break
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
		break
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
		break
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
		break
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
		break
	case "trace":
		Logger.SetLevel(logrus.TraceLevel)
		break
	default:
		Logger.SetLevel(logrus.DebugLevel)
		break
	}
	Logger.AddHook(logConf.newLfsLook())
	//定位行号
	Logger.SetReportCaller(true)
	fileAndStdoutWriter := io.MultiWriter(os.Stdout)
	Logger.SetOutput(fileAndStdoutWriter)
	MyLog = Logger.WithFields(logrus.Fields{})
}

func (logConf *loggerConf) newLfsLook() *lfshook.LfsHook {
	// 创建所有必需的目录
	conf := logConf.Logger
	logPath := conf.LogPath
	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Log file log initialization failed:%s", err.Error()))
	}
	writerLog, err := rotatelogs.New(
		logPath+"/log.%Y%m%d.log",
		//rotatelogs.WithLinkName(conf.LogPath),     // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(conf.MaxAge*conf.RotationTime)*time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(conf.RotationTime)*time.Hour),       // 日志切割时间间隔
	)
	if err != nil {
		panic(fmt.Sprintf("Log file log.log initialization failed:%s", err.Error()))
	}
	writerErr, err := rotatelogs.New(
		logPath+"/err.%Y%m%d.log",
		//rotatelogs.WithLinkName(conf.ErrPath),     // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(conf.MaxAge*conf.RotationTime)*time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(conf.RotationTime)*time.Hour),       // 日志切割时间间隔
	)
	if err != nil {
		panic(fmt.Sprintf("Log file error.log initialization failed:%s", err.Error()))
	}
	//这个同文件里，key value来标明
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		// 为不同级别设置不同的输出目的
		logrus.PanicLevel: writerErr,
		logrus.FatalLevel: writerErr,
		logrus.ErrorLevel: writerErr,
		logrus.WarnLevel:  writerLog,
		logrus.InfoLevel:  writerLog,
		logrus.DebugLevel: writerLog,
		logrus.TraceLevel: writerLog,
	}, &logrus.TextFormatter{
		ForceColors:     false, // 强制启用彩色输出
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return lfsHook
}

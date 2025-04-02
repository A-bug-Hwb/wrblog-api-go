package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"wrblog-api-go/config"
	"wrblog-api-go/router"
)

func main() {
	//设置模式
	gin.SetMode(config.Conf.Server.Mode)
	r := router.InitRouter()
	port := fmt.Sprintf(":%s", strconv.Itoa(config.Conf.Server.Port))
	//var ser *http.Server
	if config.Conf.Server.Https {
		r.RunTLS(port, "", "")
	} else {
		r.Run(fmt.Sprintf(":%s", strconv.Itoa(config.Conf.Server.Port)))
		//ser = &http.Server{
		//	Addr:    port,
		//	Handler: r,
		//}
	}
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//mylog.MyLog.Info(fmt.Sprintf("wrblog-api 启动：端口 %s", strconv.Itoa(config.Conf.Server.Port)))
	////开一个线程执行server启动操作
	//go func() {
	//	if err := ser.ListenAndServe(); err != nil {
	//		mylog.MyLog.Panic(fmt.Sprintf("wrblog-api 程序停止：%s", err))
	//		cancel()
	//	}
	//}()
	//go func() {
	//	//监听拦截关闭事件
	//	killSignal := make(chan os.Signal, 1)
	//	signal.Notify(killSignal, os.Interrupt)
	//	<-killSignal
	//	cancel()
	//}()
	//<-ctx.Done()
	//mylog.MyLog.Info("wrblog-api shutdown!")
	//ser.Shutdown(context.Background())
}

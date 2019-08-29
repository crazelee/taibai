package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"taibai/common/conf"
	"taibai/common/log"
	"taibai/router"
)

func main() {
	confPath := flag.String("conf", "conf/app.toml", "config file's path")
	flag.Parse()

	// 初始化配置
	if err := conf.Init(*confPath); err != nil {
		panic("init config failed! err: " + err.Error())
	}
	config := conf.Get()

	// 初始化log
	if err := log.Init(); err != nil {
		panic("init log failed! err: " + err.Error())
	}

	// 记录到文件。
	//f, _ := os.Create(config.Log.Path)
	//gin.DefaultWriter = io.MultiWriter(f)

	r := router.GetHttpServer()
	err := r.Run(config.HTTP.Addr)
	if err != nil {
		fmt.Println(gin.DefaultWriter, "start failed: ", err)
	}
}

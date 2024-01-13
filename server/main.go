package main

import (
	"HttpTest-Server/global"
	"HttpTest-Server/initialization"
	"HttpTest-Server/router"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	// 初始化
	initialization.HTInit()

	// 路由初始化
	r := router.Init()

	// 运行接口服务
	err := r.Run(fmt.Sprintf(":%d", global.SysPort))
	if err != nil {
		global.APP_LOG.Panic("接口服务启动失败", zap.Error(err))
	}
}

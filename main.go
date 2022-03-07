package main

import (
	"registry-manager/bootstrap"
	"registry-manager/router"
)

func main() {
	// 解压静态文件
	bootstrap.Eject()
	bootstrap.InitStatic()
	// 初始化路由
	engine := router.InitRouter()
	if err := engine.Run(":8081"); err != nil {

	}
}

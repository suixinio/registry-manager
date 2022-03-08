package main

import (
	"embed"
	"registry-manager/bootstrap"
	"registry-manager/router"
)

//go:embed assets/build
var StaticAsset embed.FS

func main() {
	// 解压静态文件
	bootstrap.StaticEmbed = StaticAsset
	bootstrap.InitStatic()
	// 初始化路由
	engine := router.InitRouter()
	if err := engine.Run(":8081"); err != nil {

	}
}

package main

import (
	"embed"
	"flag"
	"registry-manager/bootstrap"
	"registry-manager/pkg/conf"
	"registry-manager/pkg/util"
	"registry-manager/router"
)

//go:embed assets/build
var StaticAsset embed.FS

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "c", util.RelativePath("conf.ini"), "配置文件路径")
	flag.Parse()

	bootstrap.StaticEmbed = StaticAsset
	bootstrap.InitStatic()
	conf.Init(confPath)
	bootstrap.Init()
}

func main() {
	// 解压静态文件
	// 初始化路由
	engine := router.InitRouter()
	if err := engine.Run(":8081"); err != nil {

	}
}

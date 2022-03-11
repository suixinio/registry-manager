package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"registry-manager/api"
	"registry-manager/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// api router
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/"})))
	// Logger
	r.Use(middleware.Logger())
	// 静态文件
	r.Use(middleware.FrontendFileHandler())

	{
		// registry V2 的路由
		registryV2 := r.Group("/v2")
		registryV2.Use(middleware.BasicAuth())
		registryV2.Any("/*any", api.RegistryV2Any)
	}

	return r
}

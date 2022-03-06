package router

import (
	"github.com/gin-gonic/gin"
	"registry-manager/api"
	"registry-manager/middleware"
)

func InitRouter() {
	r := gin.New()

	registryV2 := r.Group("/v2")
	registryV2.Use(middleware.BasicAuth())

	registryV2.Any("/*any", api.RegistryV2Any)

	err := r.Run(":3001")
	if err != nil {
		return
	}
}

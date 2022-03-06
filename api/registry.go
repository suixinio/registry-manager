package api

import (
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
)

func RegistryV2Any(c *gin.Context) {
	var proxyUrl = new(url.URL)
	proxyUrl.Scheme = "http"
	proxyUrl.Host = "127.0.0.1:3002"
	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
	proxy.ServeHTTP(c.Writer, c.Request)
}

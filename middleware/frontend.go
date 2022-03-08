package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"registry-manager/bootstrap"
	"strings"
)

// FrontendFileHandler 前端静态文件处理
func FrontendFileHandler() gin.HandlerFunc {
	ignoreFunc := func(c *gin.Context) {
		c.Next()
	}

	if bootstrap.StaticFS == nil {
		return ignoreFunc
	}

	// 读取index.html
	file, err := bootstrap.StaticFS.Open("/index.html")
	if err != nil {
		//util.Log().Warning("静态文件[index.html]不存在，可能会影响首页展示")
		return ignoreFunc
	}

	fileContentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		//util.Log().Warning("静态文件[index.html]读取失败，可能会影响首页展示")
		return ignoreFunc
	}
	fileContent := string(fileContentBytes)

	fileServer := http.FileServer(bootstrap.StaticFS)
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		// API 跳过
		if strings.HasPrefix(path, "/api") || strings.HasPrefix(path, "/custom") || strings.HasPrefix(path, "/v2") {
			c.Next()
			return
		}
		fmt.Println(path)

		// 不存在的路径和index.html均返回index.html
		if (path == "/index.html") || (path == "/") || !bootstrap.StaticFS.Exists("/", path) {
			c.Header("Content-Type", "text/html")
			c.String(200, fileContent)
			c.Abort()
			return
		}

		// 存在的静态文件
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}

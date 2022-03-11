package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
	"registry-manager/pkg/util"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		//hostName
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}

		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		util.ALog.Info().Int("spendTime", int(stopTime.Milliseconds())).
			Str("hostName", hostName).Int("statusCode", statusCode).
			Str("ip", clientIp).Str("user-agent", userAgent).Int("dataSize", dataSize).
			Str("method", method).Str("path", path).Send()
	}
}

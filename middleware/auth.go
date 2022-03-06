package middleware

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const AuthUserKey = "user"

func SearchCredential(authValue string) (string, bool) {

	if authValue == authorizationHeader("test", "test") {
		return "test", true
	}
	return "", false
}

func BasicAuth() gin.HandlerFunc {
	realm := `Basic realm="basic-realm"`
	return func(c *gin.Context) {
		user, found := SearchCredential(c.GetHeader("Authorization"))
		fmt.Println(user, found)
		if !found {
			//	not found
			c.Header("WWW-Authenticate", realm)
			c.Header("Docker-Distribution-Api-Version", "registry/2.0")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set(AuthUserKey, user)
		c.Next()
	}
}

func authorizationHeader(user, password string) string {
	base := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(base))
}

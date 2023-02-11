package jwt

import (
	"net/http"
	"time"

	"github.com/Lyianu/wen/pkg/e"
	"github.com/Lyianu/wen/util"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		hToken := c.Request.Header.Get("Authorization")

		if hToken == "" || len(hToken) < 8 {
			code = e.INVALID_PARAMS
		} else {
			token := hToken[7:]
			code = e.SUCCESS
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

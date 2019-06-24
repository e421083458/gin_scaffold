package middleware

import (
	"errors"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
)

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range lib.GetStringSliceConf("base.http.allow_ip") {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched{
			ResponseError(c, InternalErrorCode, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}

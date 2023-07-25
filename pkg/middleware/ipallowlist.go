package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

var ipNotAllowList = []string{}

var limit = 1000
var requestNumber = map[string]int{}

func registerIpAllowListHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.ClientIP()
		for _, v := range ipNotAllowList {
			if uri == v {
				c.AbortWithStatusJSON(403, gin.H{
					"message": "Forbidden",
				})
				return
			}
		}
		ip := c.ClientIP()
		if requestNumber[ip] > limit {
			c.AbortWithStatusJSON(429, gin.H{
				"message": "Too Many Requests",
			})
			return
		}
		requestNumber[ip]++
		//clear requestNumber every 1 hour
		go func() {
			time.Sleep(1 * time.Hour)
			requestNumber = map[string]int{}
		}()
		c.Next()
	}
}

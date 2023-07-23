package middleware

import (
	v1 "FlyDB-Web/pkg/cache"
	"github.com/gin-gonic/gin"
	"strings"
)

// TODO we should add read url list to config file
var readList = []string{
	"string/get",
}

func registerQuotaHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// listen whether the current operation is a read state,
		// and the non-read state needs to modify the memory usage
		path := c.Request.URL.Path
		for _, s := range readList {
			if strings.Contains(path, s) {
				c.Next()
			} else {
				id, exists := c.Get("id")
				if !exists {
					c.AbortWithStatusJSON(401, gin.H{
						"message": "Unauthorized",
					})
					return
				}
				store := v1.GetUserStore()
				user := store.GetUser(id.(string))
				if user.UsedMem > user.TotalMem {
					c.AbortWithStatusJSON(500, gin.H{
						"message": "quota exceeded",
					})
					return
				}
				err := store.UpdateUser(user)
				if err != nil {
					c.AbortWithStatusJSON(500, gin.H{
						"message": err.Error(),
					})
					return
				}
				c.Next()
			}
		}
	}
}

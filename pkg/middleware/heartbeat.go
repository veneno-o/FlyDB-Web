package middleware

import (
	v1 "FlyDB-Web/pkg/cache"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

var heartbeat = make(map[string]time.Time)

func registerHeartBeatHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get("id")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		heartbeat[id.(string)] = time.Now()
		once := sync.Once{}
		once.Do(func() {
			go func() {
				for {
					time.Sleep(5 * time.Second)
					for k, v := range heartbeat {
						// if user has no heartbeat in 60 Minute, we will close the connection
						if time.Now().Sub(v) > 60*time.Minute {
							delete(heartbeat, k)
							user := v1.GetUserStore().GetUser(k)
							err := user.Close()
							if err != nil {
								panic(err)
							}
						}
					}
				}
			}()
		})
	}
}

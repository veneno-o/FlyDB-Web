package service

import (
	R "FlyDB-Web/pkg/result"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	R.Success(c, "ping")
}

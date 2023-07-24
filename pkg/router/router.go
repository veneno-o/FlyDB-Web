package router

import (
	"FlyDB-Web/pkg/service"
	"github.com/gin-gonic/gin"
)

// Wrapper is a wrapper of gin.Engine
func Wrapper(r *gin.Engine) {
	//define router
	r.GET("/ping", service.Ping)
	r.POST("/api/v1/string/put", service.PutString)
	r.POST("/api/v1/string/get", service.GetString)
	r.DELETE("/api/v1/string/delete", service.DeleteString)
}

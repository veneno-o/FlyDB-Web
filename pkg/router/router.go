package router

import (
	"FlyDB-Web/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Wrapper is a wrapper of gin.Engine
func Wrapper(r *gin.Engine) {
	//define router
	r.GET("/ping", service.Ping)
	r.POST("/api/v1/string/put", service.PutString)
	r.POST("/api/v1/string/get", service.GetString)
	r.DELETE("/api/v1/string/delete", service.DeleteString)
	r.POST("/api/v1/hash/hget", service.GetHash)
	r.POST("/api/v1/hash/hset", service.SetHash)
	r.DELETE("/api/v1/hash/hdel", service.DelHash)
	r.GET("/auth/github", service.AuthGithub)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

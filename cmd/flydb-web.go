package main

import (
	"FlyDB-Web/pkg/handler"
	"FlyDB-Web/pkg/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	for _, handlerFunc := range handler.GetAllHandler() {
		r.Use(handlerFunc)
	}
	router.Wrapper(r)
	err := r.Run(":1313")
	if err != nil {
		_ = fmt.Errorf("gin run error: %v", err)
	}
}

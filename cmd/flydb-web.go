package main

import (
	v1 "FlyDB-Web/pkg/cache"
	"FlyDB-Web/pkg/middleware"
	"FlyDB-Web/pkg/router"
	"fmt"
	"github.com/ByteStorage/FlyDB/structure"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := gin.Default()
	for _, handlerFunc := range middleware.GetAllMiddleware() {
		r.Use(handlerFunc)
	}
	router.Wrapper(r)
	err := r.Run(":1313")
	if err != nil {
		_ = fmt.Errorf("gin run error: %v", err)
	}
	//graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL)

	<-quit
	for _, user := range v1.GetUserStore().GetUsers() {
		for _, db := range user.Db {
			switch v := db.(type) {
			case *structure.StringStructure:
				err := v.Stop()
				if err != nil {
					fmt.Println(err)
				}
			default:
				fmt.Println("unknown type")
			}
		}
	}

}

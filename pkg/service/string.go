package service

import (
	R "FlyDB-Web/pkg/result"
	"fmt"
	"github.com/ByteStorage/FlyDB/config"
	"github.com/ByteStorage/FlyDB/structure"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func PutString(c *gin.Context) {
	fmt.Println("PutString")
	id, exists := c.Get("id")
	if !exists {
		R.Error(c, 400, "user not login")
		return
	}
	defaultOptions := config.DefaultOptions
	defaultOptions.DirPath = "/var/lib/flydb/" + id.(string)
	dbs, err := structure.NewStringStructure(defaultOptions)
	if err != nil {
		R.Error(c, 500, err.Error())
		return
	}
	var req RequestBody
	err = c.BindJSON(&req)
	if err != nil {
		R.Error(c, 400, err.Error())
		return
	}
	err = dbs.Set(req.Key, req.Value, 0)
	if err != nil {
		R.Error(c, 500, err.Error())
		return
	}
	R.Success(c, gin.H{
		"key":   req.Key,
		"value": req.Value,
	})

}

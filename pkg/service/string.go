package service

import (
	v1 "FlyDB-Web/pkg/cache"
	R "FlyDB-Web/pkg/result"
	"github.com/ByteStorage/FlyDB/structure"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func PutString(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		R.Error(c, 400, "user not login")
		return
	}
	store := v1.GetUserStore()
	user := store.GetUser(id.(string))
	for _, db := range user.Db {
		if dbs, ok := db.(*structure.StringStructure); ok {
			var req RequestBody
			err := c.BindJSON(&req)
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
			return
		}
	}
	R.Error(c, 500, "unknown type")
}

func GetString(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		R.Error(c, 400, "user not login")
		return
	}
	store := v1.GetUserStore()
	user := store.GetUser(id.(string))
	for _, db := range user.Db {
		if dbs, ok := db.(*structure.StringStructure); ok {
			var req RequestBody
			err := c.BindJSON(&req)
			if err != nil {
				R.Error(c, 400, err.Error())
				return
			}
			value, err := dbs.Get(req.Key)
			if err != nil {
				R.Error(c, 500, err.Error())
				return
			}
			R.Success(c, gin.H{
				"key":   req.Key,
				"value": value,
			})
			return
		}
	}
	R.Error(c, 500, "unknown type")
}

func DeleteString(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		R.Error(c, 400, "user not logged in")
		return
	}
	store := v1.GetUserStore()
	user := store.GetUser(id.(string))
	for _, db := range user.Db {
		if dbs, ok := db.(*structure.StringStructure); ok {
			var req RequestBody
			err := c.BindJSON(&req)
			if err != nil {
				R.Error(c, 400, err.Error())
				return
			}
			err = dbs.Del(req.Key)
			if err != nil {
				R.Error(c, 500, err.Error())
				return
			}
			R.Success(c, gin.H{
				"key": req.Key,
			})
			return
		}
	}
	R.Error(c, 500, "unknown type")
}

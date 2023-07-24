package service

import (
	v1 "FlyDB-Web/pkg/cache"
	R "FlyDB-Web/pkg/result"
	"github.com/ByteStorage/FlyDB/structure"
	"github.com/gin-gonic/gin"
)

type HashRequestBody struct {
	Key   string      `json:"key"`
	Field interface{} `json:"field"`
	Value interface{} `json:"value"`
}

func GetHash(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		R.Error(c, 400, "user not logged in")
		return
	}
	store := v1.GetUserStore()
	user := store.GetUser(id.(string))
	for _, db := range user.Db {
		if dbs, ok := db.(*structure.HashStructure); ok {
			var req HashRequestBody
			err := c.BindJSON(&req)
			if err != nil {
				R.Error(c, 400, err.Error())
				return
			}

			value, err := dbs.HGet(req.Key, req.Field)
			if err != nil {
				R.Error(c, 500, err.Error())
				return
			}

			R.Success(c, gin.H{
				"code":    "200",
				"message": "HGet success",
				"data": gin.H{
					"key":   req.Key,
					"field": req.Field,
					"value": value,
				},
			})
			return
		}
	}
	R.Error(c, 500, "unknown type")
}

func SetHash(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		R.Error(c, 400, "user not logged in")
		return
	}
	store := v1.GetUserStore()
	user := store.GetUser(id.(string))
	for _, db := range user.Db {
		if dbs, ok := db.(*structure.HashStructure); ok {
			var req HashRequestBody
			err := c.BindJSON(&req)
			if err != nil {
				R.Error(c, 400, err.Error())
				return
			}

			_, err = dbs.HSet(req.Key, req.Field, req.Value)
			if err != nil {
				R.Error(c, 500, err.Error())
				return
			}

			R.Success(c, gin.H{
				"key":   req.Key,
				"Field": req.Field,
				"value": req.Value,
			})
			return
		}
	}
	R.Error(c, 500, "unknown type")
}

func DelHash(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		R.Error(c, 400, "user not logged in")
		return
	}
	store := v1.GetUserStore()
	user := store.GetUser(id.(string))
	for _, db := range user.Db {
		if dbs, ok := db.(*structure.HashStructure); ok {
			var req HashRequestBody
			err := c.BindJSON(&req)
			if err != nil {
				R.Error(c, 400, err.Error())
				return
			}

			_, err = dbs.HDel(req.Key, req.Field)
			if err != nil {
				R.Error(c, 500, err.Error())
				return
			}

			R.Success(c, gin.H{
				"key":   req.Key,
				"Field": req.Field,
			})
			return
		}
	}
	R.Error(c, 500, "unknown type")
}

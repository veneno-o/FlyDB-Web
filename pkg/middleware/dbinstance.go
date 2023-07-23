package middleware

import (
	v1 "FlyDB-Web/pkg/cache"
	"github.com/ByteStorage/FlyDB/config"
	"github.com/ByteStorage/FlyDB/structure"
	"github.com/gin-gonic/gin"
)

func registerDbInstanceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		store := v1.GetUserStore()
		id, exists := c.Get("id")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		dbInstance, err := createAllDbInstance(id.(string))
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		err = store.CreateUser(v1.User{
			Id:       id.(string),
			Db:       dbInstance,
			TotalMem: 1024 * 1024 * 1024,
			UsedMem:  0,
		})
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
}

func createAllDbInstance(id string) ([]interface{}, error) {
	options := config.DefaultOptions
	options.DirPath = "/var/lib/flydb/" + id
	dbList := make([]interface{}, 0)
	stringStructure, err := structure.NewStringStructure(options)
	if err != nil {
		return nil, err
	}
	listStructure, err := structure.NewListStructure(options)
	if err != nil {
		return nil, err
	}
	hashStructure, err := structure.NewHashStructure(options)
	if err != nil {
		return nil, err
	}
	dbList = append(dbList, stringStructure)
	dbList = append(dbList, listStructure)
	dbList = append(dbList, hashStructure)
	return dbList, nil
}

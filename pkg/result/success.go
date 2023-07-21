package R

import "github.com/gin-gonic/gin"

// Success returns success message
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

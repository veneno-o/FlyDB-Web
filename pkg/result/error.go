package R

import "github.com/gin-gonic/gin"

// Error returns error message
func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": msg,
	})
}

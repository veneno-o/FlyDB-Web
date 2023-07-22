package middleware

import "github.com/gin-gonic/gin"

// GetAllMiddleware returns all middleware handlers
func GetAllMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		registerAuthHandler(),
	}
}

// registerAuthHandler is a middleware that checks whether the user is logged in
func registerAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// mock login
		c.Set("id", "test")
		c.Next()
	}
}

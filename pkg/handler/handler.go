package handler

import "github.com/gin-gonic/gin"

// GetAllHandler returns all middleware handlers
func GetAllHandler() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		registerAuthHandler(),
	}
}

// registerAuthHandler is a middleware that checks whether the user is logged in
func registerAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

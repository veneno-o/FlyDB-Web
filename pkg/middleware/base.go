package middleware

import "github.com/gin-gonic/gin"

// GetAllMiddleware returns all middleware handlers
func GetAllMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		registerAuthHandler(),
		registerIpAllowListHandler(),
		registerQuotaHandler(),
		registerDbInstanceHandler(),
	}
}

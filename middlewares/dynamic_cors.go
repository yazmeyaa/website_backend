package middlewares

import (
	"net/http"
	"strings"
	"yazmeyaa_projects/service"

	"github.com/gin-gonic/gin"
)

func DynamicCorsMiddleware(corsService service.CorsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		if origin != "" {
			record, err := corsService.GetRecord(c, origin)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "CORS policy does not allow this origin"})
				return
			}
			if !record.OriginAllowed {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", strings.Join(record.AllowedMethods, ", "))
			c.Header("Access-Control-Allow-Headers", strings.Join(record.AllowedHeaders, ", "))

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusOK)
				return
			}
		}

		c.Next()
	}
}

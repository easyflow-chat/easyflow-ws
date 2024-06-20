package middleware

import (
	"easyflow-ws/src/net"

	"github.com/gin-gonic/gin"
)

func InjectSup(h *net.Supervisor) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("super", h)
		c.Next()
	}
}

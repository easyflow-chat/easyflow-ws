package middleware

import (
	"easyflow-ws/src/common"
	"github.com/gin-gonic/gin"
)

func InjectCfg(cfg *common.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("cfg", cfg)
		c.Next()
	}
}

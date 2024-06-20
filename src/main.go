package main

import (
	"easyflow-ws/src/api"
	"easyflow-ws/src/common"
	"easyflow-ws/src/net"
	"os"

	"github.com/gin-gonic/gin"
)

func injectHub(h *net.Supervisor) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("super", h)
		c.Next()
	}
}
func main() {
	config := common.LoadDefaultConfig()
	logger := common.NewLogger(os.Stdout, "main")
	supervisor := net.NewSupervisor()

	router := gin.Default()
	router.Use(injectHub(supervisor))
	router.GET("/ws", api.WebsocketListener)

	logger.PrintfInfo("Starting ws-worker on port %s", config.Port)
	router.Run(":" + config.Port)
}

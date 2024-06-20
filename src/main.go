package main

import (
	"easyflow-ws/src/api"
	"easyflow-ws/src/common"
	"easyflow-ws/src/middleware"
	"easyflow-ws/src/net"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config := common.LoadDefaultConfig()
	logger := common.NewLogger(os.Stdout, "main")
	supervisor := net.NewSupervisor()

	router := gin.Default()
	router.Use(middleware.InjectSup(supervisor))
	router.Use(middleware.InjectCfg(config))
	router.GET("/ws", api.WebsocketListener)

	logger.PrintfInfo("Starting ws-worker on port %s", config.Port)
	router.Run(":" + config.Port)
}

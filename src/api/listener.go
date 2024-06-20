package api

import (
	"easyflow-ws/src/common"
	"easyflow-ws/src/net"
	"os"

	"github.com/gin-gonic/gin"
)

func WebsocketListener(c *gin.Context) {
	raw_hub, ok := c.Get("hub")
	logger := common.NewLogger(os.Stdout, "WsHandler")
	if !ok {
		logger.PrintfError("Failed to retrieve Hub. Exiting")
		os.Exit(1)
	}
	hub := raw_hub.(net.Hub)

	
}

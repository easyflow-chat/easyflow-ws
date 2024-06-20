package api

import (
	"easyflow-ws/src/common"
	"easyflow-ws/src/net"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"os"
)

var upgrader = websocket.Upgrader{}
var logger = common.NewLogger(os.Stdout, "WebsocketListener")

func WebsocketListener(c *gin.Context) {
	raw_hub, ok := c.Get("hub")
	if !ok {
		logger.PrintfError("Failed to retrieve Hub. Exiting")
		panic("")
	}
	hub, ok := raw_hub.(*net.Hub)
	if !ok {
		logger.PrintfError("Failed to convert Hub. Exiting")
		os.Exit(1)
	}

	req := c.Request
	res := c.Writer
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		logger.PrintfError("An error occured while trying to upgrade to websocket: %v", err)
		panic("")
	}

	userId := c.Query("userId")
	logger.PrintfInfo("%s", userId)
	if userId == "" {
		logger.PrintfError("No userid was provided")
		panic("")
	}

	info := net.ClientInfo{
		Uid: userId,
	}
	client := net.NewClient(conn, &info)
	logger.Printf("hal")

	hub.Insert(client)
	// In your WebsocketListener function
	logger.PrintfInfo("Calling WebsocketHandler for user %s", userId)
	WebsocketHandler(client) // Check if you meant to call this asynchronously with 'go'

}

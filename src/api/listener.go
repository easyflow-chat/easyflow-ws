package api

import (
	"easyflow-ws/src/common"
	"easyflow-ws/src/net"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var logger = common.NewLogger(os.Stdout, "WebsocketListener")

func WebsocketListener(c *gin.Context) {
	raw_sup, ok := c.Get("super")
	if !ok {
		logger.PrintfError("Failed to retrieve Supervisor. Exiting")
		c.AbortWithStatusJSON(http.StatusInternalServerError, ApiError{
			ErrorCode: http.StatusInternalServerError,
			Details:   "Failed to retrieve Supervisor from context",
		})
		return
	}
	sup, ok := raw_sup.(*net.Supervisor)
	if !ok {
		logger.PrintfError("Failed to convert Supervisor. Exiting")
		c.AbortWithStatusJSON(http.StatusInternalServerError, ApiError{
			ErrorCode: http.StatusInternalServerError,
			Details:   "Failed to convert Supervisor",
		})
		return
	}

	raw_cfg, ok := c.Get("cfg")
	if !ok {
		logger.PrintfError("Failed to retrieve Config. Exiting")
		c.AbortWithStatusJSON(http.StatusInternalServerError, ApiError{
			ErrorCode: http.StatusInternalServerError,
			Details:   "Failed to retrieve Config from context",
		})
		return
	}

	cfg, ok := raw_cfg.(*common.Config)
	if !ok {
		logger.PrintfError("Failed to convert Config. Exiting")
		c.AbortWithStatusJSON(http.StatusInternalServerError, ApiError{
			ErrorCode: http.StatusInternalServerError,
			Details:   "Failed to convert Config",
		})
		return
	}

	userId := c.Query("userId")
	if userId == "" {
		logger.PrintfError("No userid was provided")
		c.JSON(http.StatusInternalServerError, ApiError{
			ErrorCode: http.StatusBadRequest,
			Details:   "No userid was provided",
		})
		return
	}

	info := net.ClientInfo{
		Uid:      userId,
		SocketId: uuid.New().String(),
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.PrintfError("WebSocket upgrade failed: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ApiError{
			ErrorCode: http.StatusInternalServerError,
			Details:   fmt.Errorf("WebSocket upgrade failed: %v", err),
		})
		return
	}

	go func() {
		client := net.NewClient(conn, &info)
		defer client.Close()
		sup.Insert(client)
		defer sup.Remove(client)

		logger.PrintfInfo("Accepted user with SocketId: %s", client.Info.SocketId)
		logger.PrintfInfo("Active connections: %d", len(sup.Clients))
		err = WebsocketHandler(client, time.Duration(cfg.Timeout))
		logger.Printf("WebsocketHandler returned: %v", err)
		if err != nil {
			logger.PrintfError("An error occurred while handling websocket: %v", err)
		}
		logger.PrintfInfo("Closed connection with: %s", client.Info.SocketId)
	}()
}

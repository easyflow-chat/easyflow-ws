package main

import (
	"easyflow-ws/src/common"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config := common.LoadDefaultConfig()
	logger := common.NewLogger(os.Stdout, "main")

	router := gin.Default()

	logger.PrintfInfo("Starting ws-worker on port %s", config.Port)
	router.Run(":" + config.Port)
}

package main

import (
	"easyflow-ws/src/common"
	"os"
)

func main() {
	config := common.LoadDefaultConfig()
	logger := common.NewLogger(os.Stdout, "main")
	logger.Printf("%s", config)
}

package api

import (
	"easyflow-ws/src/common"
	"easyflow-ws/src/net"
	"os"
)

var wsLogger = common.NewLogger(os.Stdout, "WebsocketHandler")

func WebsocketHandler(client *net.Client) {
	wsLogger.PrintfInfo("Accepted user with Id: %s", client.Info)
	for {
		err := client.Read()
		if err != nil {
			wsLogger.PrintfError("While reading from user an error occured: %v", err)
			client.Close()
			break
		}
		client.OutBuffer <- client.InBuffer
		err = client.Send()
		if err != nil {
			wsLogger.PrintfError("While sending to user an error occured: %v", err)
			client.Close()
			break
		}
	}
	wsLogger.PrintfInfo("Closed connection with: %s", client.Info.Uid)
}

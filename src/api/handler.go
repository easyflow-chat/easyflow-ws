package api

import (
	"easyflow-ws/src/common"
	"easyflow-ws/src/net"
	"errors"
	"os"
	"time"
)

var wsLogger = common.NewLogger(os.Stdout, "WebsocketHandler")

func WebsocketHandler(client *net.Client) error {
	timeoutDuration := 5 * time.Second
	timer := time.NewTimer(timeoutDuration)
	defer timer.Stop() // Ensure the timer is cleaned up properly

	for {
		// Set up the timer for each loop iteration
		if !timer.Stop() {
			<-timer.C
		}
		timer.Reset(timeoutDuration)

		// Blocking read using a goroutine to handle the timeout
		msgChan := make(chan error, 1)
		go func() {
			err := client.Read()
			msgChan <- err
		}()

		select {
		case err := <-msgChan:
			if err != nil {
				wsLogger.PrintfError("While reading from user an error occurred: %v", err)
				client.Close()
				return err
			}
			// Assuming successful read, process and send response
			client.OutBuffer <- client.InBuffer
			if err := client.Send(); err != nil {
				wsLogger.PrintfError("While sending to user an error occurred: %v", err)
				client.Close()
				return err
			}

		case <-timer.C:
			wsLogger.PrintfError("Timeout occurred while waiting for messages from user: %s", client.Info.SocketId)
			client.Close()
			return errors.New("timeout waiting for messages")
		}
	}
}

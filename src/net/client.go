package net

import (
	"easyflow-ws/src/common"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var tempLogger = common.NewLogger(os.Stdout, "TempLogger")

const (
	ReadDeadline = 4 * time.Second
)

type ClientInfo struct {
	Uid string
}

type Client struct {
	Info      *ClientInfo
	conn      *websocket.Conn
	OutBuffer chan *common.Vector[byte]
	InBuffer  *common.Vector[byte]
}

func NewClient(conn *websocket.Conn, info *ClientInfo) *Client {
	c := Client{
		Info:      info,
		conn:      conn,
		OutBuffer: make(chan *common.Vector[byte], 10),
		InBuffer:  common.NewVector[byte](),
	}
	return &c
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Send() error {
	val := <-c.OutBuffer
	tempLogger.PrintfInfo("Sending message: %s", val.Devectorize())
	return c.conn.WriteMessage(websocket.TextMessage, val.Devectorize())
}

func (c *Client) Read() error {
	_, msg, err := c.conn.ReadMessage()
	if err != nil {
		return err
	}
	c.InBuffer = common.Vectorize(msg)
	return nil
}

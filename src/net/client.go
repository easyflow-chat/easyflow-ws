package net

import (
	"easyflow-ws/src/common"
	"time"

	"github.com/gorilla/websocket"
)

const (
	ReadDeadline = 4 * time.Second
)

type ClientInfo struct {
	Uid string
}

type Client struct {
	info       *ClientInfo
	conn       *websocket.Conn
	out_buffer chan *common.Vector[byte]
	in_buffer  *common.Vector[byte]
}

func NewClient(conn *websocket.Conn, info *ClientInfo) *Client {
	c := Client{
		info: info,
		conn: conn,
	}
	c.out_buffer <- common.NewVector[byte]()
	c.in_buffer = common.NewVector[byte]()
	return &c
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Send() error {
	val := <-c.out_buffer
	return c.conn.WriteMessage(websocket.TextMessage, val.Devectorize())
}

func (c *Client) Read() error {
	_, msg, err := c.conn.ReadMessage()
	if err != nil {
		return err
	}
	c.in_buffer = common.Vectorize(msg)
	return nil
}

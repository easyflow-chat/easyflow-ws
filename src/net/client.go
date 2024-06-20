package net

import (
	"easyflow-ws/src/common"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn       *websocket.Conn
	out_buffer chan *common.Vector[byte]
	in_buffer  *common.Vector[byte]
}

func NewClient(conn *websocket.Conn) *Client {
	c := Client{
		conn: conn,
	}
	c.out_buffer <- common.NewVector[byte]()
	c.in_buffer = common.NewVector[byte]()
	return &c
}

func (c *Client) Close() {
	c.conn.Close()
}

package net

import (
	"easyflow-ws/src/common"
	"log"
)

type Supervisor struct {
	Clients map[string]*Client
}

func NewSupervisor() *Supervisor {
	return &Supervisor{
		Clients: make(map[string]*Client),
	}
}

func (h *Supervisor) Insert(client *Client) {
	h.Clients[client.Info.SocketId] = client
	log.Printf("Active connections: %d", len(h.Clients))
}

func (h *Supervisor) Remove(client *Client) {
	delete(h.Clients, client.Info.SocketId)
	log.Printf("Active connections: %d", len(h.Clients))
}

func (h *Supervisor) Broadcast(msg string) error {
	for _, client := range h.Clients {
		client.OutBuffer <- common.Vectorize([]byte(msg))
	}
	return nil
}

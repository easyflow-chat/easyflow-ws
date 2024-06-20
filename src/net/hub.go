package net

import "easyflow-ws/src/common"

type Hub struct {
	next_id int
	Clients *[]Client
}

func NewHub() *Hub {
	return &Hub{
		next_id: 0,
	}
}

func (h *Hub) Broadcast(msg string) error {
	if len(*h.Clients) > 0 {
		for _, client := range *h.Clients {
			bytes := []byte(msg)
			client.out_buffer <- common.Vectorize(bytes)
			err := client.Send()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

package websocket

import "github.com/gorilla/websocket"

type Client struct {
	hub      *Hub
	id       string
	socket   *websocket.Conn
	outbound chan []byte
}

func NewClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		hub:      hub,
		id:       socket.RemoteAddr().String(),
		socket:   socket,
		outbound: make(chan []byte),
	}
}

func (c *Client) Write() {
	for message := range c.outbound {
		c.socket.WriteMessage(websocket.TextMessage, message)
	}
	c.socket.WriteMessage(websocket.CloseMessage, []byte{})
}

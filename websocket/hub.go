package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool, 0),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (h *Hub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "could not open websocket connection", http.StatusBadRequest)
	}

	client := NewClient(h, socket)
	h.register <- client

	go client.Write()

}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.OnRegister(client)
		case client := <-h.unregister:
			h.OnUnregister(client)
		}
	}
}

func (h *Hub) OnRegister(client *Client) {
	log.Printf("Client %s registered\n", client.id)

	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.clients[client] = true
}

func (h *Hub) OnUnregister(client *Client) {
	log.Printf("Client %s has unregistered\n", client.id)

	h.mutex.Lock()
	defer h.mutex.Unlock()
	delete(h.clients, client)
	client.socket.Close()
}

func (h *Hub) Broadcast(message interface{}, senderClient *Client) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}
	for client := range h.clients {
		if client != senderClient {
			client.outbound <- data
		}
	}
}

package websocket

import (
	"github.com/oklog/ulid/v2"
	"sync"
)

// Hub manages active clients, grouped by ULIDs.
type Hub struct {
	// Groups clients by ULID
	clients map[ulid.ULID]map[*Client]bool

	// Message input channel with ULID group context
	broadcast chan *Message

	// Client registration with ULID group context
	register chan *RegisterRequest

	// Client unregistration
	unregister chan *UnregisterRequest

	// Mutex for ensuring thread-safe operations
	mu sync.Mutex
}

// Message represents a message with its associated ULID group.
type Message struct {
	GroupID ulid.ULID
	Content []byte
}

// RegisterRequest represents a new client connection and its ULID group.
type RegisterRequest struct {
	Client  *Client
	GroupID ulid.ULID
}

// UnregisterRequest for removing a client from a ULID group.
type UnregisterRequest struct {
	Client  *Client
	GroupID ulid.ULID
}

// NewHub Create a new Hub
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[ulid.ULID]map[*Client]bool),
		broadcast:  make(chan *Message),
		register:   make(chan *RegisterRequest),
		unregister: make(chan *UnregisterRequest),
	}
}

// Run handles registering, unregistering, and broadcasting.
func (h *Hub) Run() {
	for {
		select {
		case req := <-h.register:
			// Add the client to the specified ULID group
			h.mu.Lock()
			if _, exists := h.clients[req.GroupID]; !exists {
				h.clients[req.GroupID] = make(map[*Client]bool)
			}
			h.clients[req.GroupID][req.Client] = true
			h.mu.Unlock()

		case req := <-h.unregister:
			// Remove the client from the specified ULID group
			h.mu.Lock()
			if group, exists := h.clients[req.GroupID]; exists {
				if _, found := group[req.Client]; found {
					delete(group, req.Client)
					close(req.Client.send)
					if len(group) == 0 {
						// Deletes empty groups
						delete(h.clients, req.GroupID)
					}
				}
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			// Send message only to clients in the specified ULID group
			h.mu.Lock()
			if group, exists := h.clients[message.GroupID]; exists {
				for client := range group {
					select {
					case client.send <- message.Content:
					default:
						close(client.send)
						delete(group, client)
					}
				}
			}
			h.mu.Unlock()
		}
	}
}

// Broadcast sends a message to the given ULID group.
func (h *Hub) Broadcast(message *Message) {
	h.broadcast <- message
}

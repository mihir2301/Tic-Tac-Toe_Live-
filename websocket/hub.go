package websock

import "sync"

type Hub struct {
	Clients map[string]*Client //playerId->Client
	mu      sync.Mutex
}

//New HUb initializes the HUB
func NewHub() *Hub {
	return &Hub{
		Clients: make(map[string]*Client),
	}
}

//Register a new Client in the Hub
func (h *Hub) RegisterNewClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.Clients[client.PlayerID] = client
}

//unregister a client
func (h *Hub) UnregisterAClient(playerID string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.Clients, playerID)
}

//Brodcast Send Message to both players in the game
func (h *Hub) Brodcast(gameID string, message interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, client := range h.Clients {
		if client.GameID == gameID {
			client.Conn.WriteJSON(message)
		}
	}
}

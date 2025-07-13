package chat

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

type Message struct {
	SenderId   int64  `json:"senderId"`
	ReceiverId int64  `json:"receiverId"`
	Type       string `json:"type"`
	Content    string `json:"content"`
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

//func (h *Hub) Run() {
//	for {
//		select {
//		case client := <-h.register:
//			h.clients[client] = true
//		case client := <-h.unregister:
//			if _, ok := h.clients[client]; ok {
//				delete(h.clients, client)
//				close(client.send)
//			}
//		case message := <-h.broadcast:
//			for client := range h.clients {
//				var msg Message
//				err := json.Unmarshal(message, &msg)
//				if err != nil {
//					log.Println("Failed to unmarshal message:", err)
//					break
//				}
//				if msg.ReceiverId == client.userId {
//					select {
//					case client.send <- message:
//
//					default:
//						close(client.send)
//						delete(h.clients, client)
//					}
//				}
//			}
//		}
//	}
//}

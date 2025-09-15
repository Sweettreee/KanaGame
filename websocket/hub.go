package websocket

import (
	"log"
	"sync"
)

type MessageData struct {
	Sender  *Connection
	Content string
}

type Room struct {
	ID         string
	clients    map[*Connection]bool
	broadcast  chan []byte
	msgQueue   chan MessageData
	register   chan *Connection
	unregister chan *Connection
	mu         sync.Mutex
}

func NewRoom(id string) *Room {
	return &Room{
		ID:         id,
		clients:    make(map[*Connection]bool),
		msgQueue:   make(chan MessageData),
		broadcast:  make(chan []byte),
		register:   make(chan *Connection),
		unregister: make(chan *Connection),
	}
}

func (r *Room) Run(s *Server) {
	go r.processMsgQueue(process)
	for {
		select {
		case conn := <-r.register:
			r.clients[conn] = true
			conn.Room = r
			log.Println("register Room: ", r.ID, len(r.clients), "are here.")
		case conn := <-r.unregister:
			if _, ok := r.clients[conn]; ok {
				log.Println("unregister Room: ", r.ID)
				delete(r.clients, conn)
				close(conn.Send)
				if s.RemoveRoomIfEmpty(r) {
					return
				}
			}
		case msg := <-r.broadcast:
			for conn := range r.clients {
				select {
				case conn.Send <- msg:
				default:
					delete(r.clients, conn)
					close(conn.Send)
				}
			}
		}
	}
}

func (r *Room) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()

	for client := range r.clients {
		close(client.Send)
		delete(r.clients, client)
	}
	close(r.broadcast)
	close(r.msgQueue)
	close(r.register)
	close(r.unregister)
}

// processMsgQueue 함수는 GameProcess쪽에서 호출해서 사용 //
// process함수는 방이름, MessageData를 받아서 가공후 boardcast할 내용을 string으로 반환
func (r *Room) processMsgQueue(process func(string, MessageData) string) {
	for msgData := range r.msgQueue {
		processedMsg := process(r.ID, msgData)
		r.broadcast <- []byte(processedMsg)
	}
	log.Println("process is over.")
}

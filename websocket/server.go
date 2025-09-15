// 일단 websocket관련 코드를 작성하면서 같은 패키지로 묶긴 했지만 라우딩을 위해 다른 디렉토리로 이동할 가능성 있음.
package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Server struct {
	Rooms map[string]*Room
	mu    sync.RWMutex
}

func NewServer() *Server {
	return &Server{
		Rooms: make(map[string]*Room),
	}
}

func (s *Server) GetOrCreateRoom(roomID string) *Room {
	s.mu.RLock()
	room, ok := s.Rooms[roomID]
	s.mu.RUnlock()

	if ok {
		return room
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if room, ok = s.Rooms[roomID]; ok {
		return room
	}

	room = NewRoom(roomID)
	s.Rooms[roomID] = room
	return room
}

func (s *Server) RemoveRoomIfEmpty(r *Room) {
	s.mu.Lock()
	defer s.mu.Unlock()

	room, ok := s.Rooms[r.ID]
	if !ok {
		return
	}

	if len(room.clients) == 0 {
		r.Close()
		delete(s.Rooms, r.ID)
	}
}

func (s *Server) HandleWebSocket(w http.ResponseWriter, r *http.Request, roomID string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}

	client := &Connection{
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	room := s.GetOrCreateRoom(roomID)
	room.register <- client

	go client.ReadLoop(func(c *Connection, msg []byte) {
		// 방 전체에 메시지 브로드캐스트
		room.broadcast <- msg
	})

	go client.WriteLoop()
}

// 일단 websocket관련 코드를 작성하면서 같은 패키지로 묶긴 했지만 라우딩을 위해 다른 디렉토리로 이동할 가능성 있음.
package websocket

import (
	"fmt"
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
	go room.Run(s)
	s.Rooms[roomID] = room
	log.Println("Create Room: ", room.ID)
	return room
}

func (s *Server) RemoveRoomIfEmpty(r *Room) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	room, ok := s.Rooms[r.ID]
	if !ok {
		return false
	}

	if len(room.clients) == 0 {
		r.Close()
		delete(s.Rooms, r.ID)
		log.Println("Close Room: ", r.ID)
		return true
	}
	return false
}

// 테스트용 임시 프로세스 Room.Run 메소스내부의 processMsgQueue메소드의 위치도 이후 변경 필요.
func process(RoomId string, MsgData MessageData) string {
	result := fmt.Sprintf("[%s]: %s", RoomId, MsgData.Content)
	log.Println(result)
	return result
}

// 테스트용 핸들러
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
		room.msgQueue <- MessageData{client, string(msg)}
	})

	go client.WriteLoop()
}

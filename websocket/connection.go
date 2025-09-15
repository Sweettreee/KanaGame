package websocket

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	Conn *websocket.Conn
	Send chan []byte
	mu   sync.Mutex
	Room *Room // 어느 방에 속했는지 추적
}

func (c *Connection) WriteMessage(messageType int, data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Conn.WriteMessage(messageType, data)
}

func (c *Connection) ReadLoop(onMessage func(*Connection, []byte)) {
	defer func() {
		if c.Room != nil {
			c.Room.unregister <- c
		}
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		onMessage(c, msg)
	}
}

func (c *Connection) WriteLoop() {
	defer c.Conn.Close()
	for msg := range c.Send {
		if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("write error:", err)
			break
		}
	}
}

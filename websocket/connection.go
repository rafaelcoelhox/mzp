package websocket

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type ClientConnection struct {
	conn *websocket.Conn
}

func NewClientConnection(conn *websocket.Conn) *ClientConnection {
	return &ClientConnection{conn: conn}
}

func (c *ClientConnection) HandleMessages() {
	defer func() {
		c.conn.Close()
		log.Println("Conexão WebSocket fechada!")
	}()

	for {
		messageType, messageContent, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Erro na leitura: %v", err)
			}
			return
		}

		response := c.processMessage(messageContent)
		if err := c.conn.WriteMessage(messageType, response); err != nil {
			log.Printf("Erro na escrita: %v", err)
			return
		}
	}
}

func (c *ClientConnection) processMessage(msg []byte) []byte {
	timeReceive := time.Now().Format(time.RFC3339)
	return []byte(fmt.Sprintf("Sua mensagem: %s. Recebida em: %s", msg, timeReceive))
}

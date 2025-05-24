package server

import (
	"github.com/gorilla/websocket"
)

func (s *serverStruct) broadcast() {
	for {
		messobj := <-messChan
		for conn := range clients {
			if conn != messobj.conn {
				if err := conn.WriteMessage(websocket.TextMessage, messobj.message); err != nil {
					s.logger.Error(err.Error())
					mutex.Lock()
					delete(clients, conn)
					mutex.Unlock()
				}
			}
		}
	}
}

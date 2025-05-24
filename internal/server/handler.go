package server

import (
	"net/http"
)

func (s serverStruct) connect(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Couldn't upgrade to websocket", http.StatusInternalServerError)
		s.logger.Error(err.Error())
	}
	mutex.Lock()
	clients[conn] = true
	s.logger.Info("conn added")
	mutex.Unlock()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			http.Error(w, "Couldn't read the request", http.StatusInternalServerError)
			s.logger.Error(err.Error())
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
		}
		messObj := messStruct{message: message, conn: conn}
		messChan <- messObj
	}
}

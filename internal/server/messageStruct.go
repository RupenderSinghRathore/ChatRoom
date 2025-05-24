package server

import "github.com/gorilla/websocket"

type messStruct struct {
	message []byte
	conn    *websocket.Conn
}

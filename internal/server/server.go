package server

import (
	"log/slog"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

type serverStruct struct {
	logger   *slog.Logger
	serve    *http.Server
	upgrader websocket.Upgrader
}

var (
	mutex    sync.Mutex
	clients  = make(map[*websocket.Conn]bool)
	messChan = make(chan messStruct)
)

func Start(port string) error {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server := serverStruct{
		logger: logger,
		upgrader: websocket.Upgrader{
			WriteBufferSize: 1024,
			ReadBufferSize:  1024,
		},
	}

	go server.broadcast()

	mux := server.newRouter()

	logger.Info("Starting started", "port", port)
	err := http.ListenAndServe(port, mux)
	return err
}

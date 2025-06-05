package main

import (
	"log"

	"github.com/RupenderSinghRathore/ChatRoom/internal/server"
)

var arg, port, user = ParseClArgs()

func main() {

	switch arg {
	case "start":
		err := server.Start(port)
		log.Fatal(err)
	case "connect":
		err := connecting(user)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("$$ Unexpected arg given $$")
	}
}

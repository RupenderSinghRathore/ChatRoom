package main

import (
	"log"

	"github.com/RupenderSinghRathore/ChatRoom/internal/server"
)

func main() {
	arg, port, user := ParseClArgs()

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

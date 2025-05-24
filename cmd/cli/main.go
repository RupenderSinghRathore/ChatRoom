package main

import (
	"log"

	"github.com/RupenderSinghRathore/ServerChatRoom/internal/server"
)

func main() {
	arg, port := ParseClArgs()
	switch arg {
	case "start":
		err := server.Start(port)
		log.Fatal(err)
	case "connect":
		err := connecting()
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("$$ Unexpected arg given $$")
	}
}

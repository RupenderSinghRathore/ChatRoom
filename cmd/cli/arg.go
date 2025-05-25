package main

import (
	"flag"
	"fmt"
	"log"
)

func ParseClArgs() (string, string, string) {
	port := flag.String("port", ":8080", "Give the port to start server")
	user := flag.String("user", "Anonymous", "Username")
	flag.Parse()

	fmt.Println("user:", *user)

	args := flag.Args()

	if len(args) == 0 {
		log.Fatalln("$$ Expected args 'start' or 'connect' none given $$")
	}
	return args[0], *port, *user
}

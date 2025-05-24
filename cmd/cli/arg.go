package main

import (
	"flag"
	"log"
)

func ParseClArgs() (string, string) {
	port := flag.String("port", ":8080", "Give the port to start server")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatalln("$$ Expected args 'start' or 'connect' none given $$")
	}
	return args[0], *port
}

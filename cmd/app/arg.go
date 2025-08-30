package main

import (
	"flag"
	"log"
	"os"
)

func ParseClArgs() (string, string, string) {
	if len(os.Args) < 2 {
		log.Fatalln("$$ Expected args 'start' or 'connect' none given $$")
	}
	arg := os.Args[1]
	if arg != "start" && arg != "connect" {
		log.Fatalf("$$ Expected args 'start' or 'connect' \"%v\" given $$", arg)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	user := flag.String("user", "Anonymous", "Username")
	err := flag.CommandLine.Parse(os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}

	return arg, port, *user
}

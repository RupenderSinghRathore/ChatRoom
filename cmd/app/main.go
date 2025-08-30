package main

import (
	"flag"
	"log"
)

func main() {
	user := flag.String("user", "Anonymous", "Set username")
	help := flag.Bool("help", false, "help message")
	flag.Parse()
	if *help {
		printHelp()
		return
	}
	err := connecting(*user)
	if err != nil {
		log.Fatal(err)
	}
}

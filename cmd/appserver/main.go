package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	port := flag.String("port", "8080", "Give the port to start server")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("$$ Expected args 'start' or 'connect' none given $$")
		os.Exit(1)
	}
	arg := args[0]
	switch arg {
	case "start":
		fmt.Println("Starting the server on port:", *port)
	case "connect":
		fmt.Println("Connecting to server")
	default:
		fmt.Println("$$ Unexpected arg given $$")
	}

	// fmt.Println("args:", arg)

	// scanner := bufio.NewScanner(os.Stdin)
	// for {
	// 	fmt.Print("> ")
	// 	scanner.Scan()
	// 	text := scanner.Text()
	// 	fmt.Println(text)
	// }
}

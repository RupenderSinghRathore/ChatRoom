package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

func connecting(user string) error {
	url := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/connect"}
	header := make(map[string][]string)
	header["user"] = []string{user}
	c, _, err := websocket.DefaultDialer.Dial(url.String(), header)
	if err != nil {
		return err
	}
	defer c.Close()

	go func() {

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					fmt.Println("bye..")
					os.Exit(0)
				} else {
					log.Fatal("err:", err)
				}
			}
			fmt.Printf("Bro said: %v", string(message))
		}
	}()

	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		if text == "" {
			continue
		}
		if err = c.WriteMessage(websocket.TextMessage, []byte(text)); err != nil {
			return err
		}
	}
}

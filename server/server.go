// Copyright (c) 2013 Joshua Elliott
// Released under the MIT License
// http://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"github.com/dmitriyGarden/turnpike"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func main() {
	s := turnpike.NewServer()
	// for Go1.1, this will be just websocket.Handler(s.HandleWebsocket)
	http.Handle("/", websocket.Handler(turnpike.HandleWebsocket(s)))
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

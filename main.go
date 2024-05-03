package main

import (
	"log"

	"github.com/tnp2004/chat-websocket/handle"
	"github.com/tnp2004/chat-websocket/server"
)

func main() {
	h := handle.NewWebsocketHandle()
	s := server.NewServer(h)

	if err := s.Listening(); err != nil {
		log.Println(err.Error())
	}
}

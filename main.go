package main

import (
	"log"

	"github.com/tnp2004/chat-websocket/config"
	"github.com/tnp2004/chat-websocket/handle"
	"github.com/tnp2004/chat-websocket/server"
)

func main() {
	conf := config.ConfigGetting()
	h := handle.NewWebsocketHandle(conf)
	s := server.NewServer(conf.Server, h)

	if err := s.Listening(); err != nil {
		log.Println(err.Error())
	}
}

package main

import (
	"github.com/tnp2004/chat-websocket/handle"
	"github.com/tnp2004/chat-websocket/server"
)

func main() {
	h := handle.NewWebsocketHandle()
	s := server.NewServer(h)

	s.Listening()
}

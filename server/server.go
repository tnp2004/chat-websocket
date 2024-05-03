package server

import (
	"net/http"

	"github.com/tnp2004/chat-websocket/exception"
	"github.com/tnp2004/chat-websocket/handle"
	"golang.org/x/net/websocket"
)

type server struct {
	handle handle.IWebsocketHandle
}

type IServer interface {
	Listening() error
}

func NewServer(handle handle.IWebsocketHandle) IServer {
	return &server{handle}
}

func (s *server) Listening() error {
	http.Handle("/ws", websocket.Handler(s.handle.IncomingHandle))

	if err := http.ListenAndServe(":3000", nil); err != nil {
		return &exception.ListeningFailed{Addr: "ws://localhost:3000"}
	}

	return nil
}

package server

import (
	"fmt"
	"net/http"

	"github.com/tnp2004/chat-websocket/config"
	"github.com/tnp2004/chat-websocket/exception"
	"github.com/tnp2004/chat-websocket/handle"
	"golang.org/x/net/websocket"
)

type server struct {
	conf   *config.Server
	handle handle.IWebsocketHandle
}

type IServer interface {
	Listening() error
}

func NewServer(conf *config.Server, handle handle.IWebsocketHandle) IServer {
	return &server{conf, handle}
}

func (s *server) Listening() error {
	http.Handle("/ws", websocket.Handler(s.handle.IncomingHandler))
	http.Handle("/rooms", websocket.Handler(s.handle.RoomHandler))
	http.Handle("/rooms/create", websocket.Handler(s.handle.CreateRoomHandler))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.conf.Port), nil); err != nil {
		Addr := fmt.Sprintf("ws://%s:%d", s.conf.Host, s.conf.Port)
		return &exception.ListeningFailed{Addr: Addr}
	}

	return nil
}

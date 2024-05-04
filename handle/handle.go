package handle

import (
	"fmt"
	"io"

	"github.com/tnp2004/chat-websocket/config"
	"golang.org/x/net/websocket"
)

type websocketHandle struct {
	conf  *config.Config
	conns map[*websocket.Conn]bool
}

type IWebsocketHandle interface {
	IncomingHandle(ws *websocket.Conn)
}

func NewWebsocketHandle(conf *config.Config) IWebsocketHandle {
	return &websocketHandle{conf: conf, conns: make(map[*websocket.Conn]bool)}
}

func (h *websocketHandle) IncomingHandle(ws *websocket.Conn) {
	fmt.Println("New connection from:", ws.RemoteAddr())

	h.conns[ws] = true

	h.readLoop(ws)
}

func (h *websocketHandle) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("Read error: ", err.Error())
			continue
		}
		msg := buf[:n]

		ws.Write([]byte(fmt.Sprintf("Message from server: %s", msg)))
	}
}

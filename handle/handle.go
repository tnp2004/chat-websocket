package handle

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

type websocketHandle struct {
	conns map[*websocket.Conn]bool
}

type IWebsocketHandle interface {
	IncomingHandle(ws *websocket.Conn)
}

func NewWebsocketHandle() IWebsocketHandle {
	return &websocketHandle{conns: make(map[*websocket.Conn]bool)}
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

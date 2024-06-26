package handle

import (
	"fmt"
	"io"
	"log"

	"github.com/tnp2004/chat-websocket/config"
	"github.com/tnp2004/chat-websocket/exception"
	"golang.org/x/net/websocket"
)

type websocketHandle struct {
	conf  *config.Config
	rooms map[string]connection
}

type connection struct {
	conns map[*websocket.Conn]bool
}

type IWebsocketHandle interface {
	IncomingHandler(ws *websocket.Conn)
	RoomHandler(ws *websocket.Conn)
	CreateRoomHandler(ws *websocket.Conn)
}

func NewWebsocketHandle(conf *config.Config) IWebsocketHandle {
	return &websocketHandle{conf: conf, rooms: make(map[string]connection)}
}

func (h *websocketHandle) IncomingHandler(ws *websocket.Conn) {
	fmt.Println("New connection from:", ws.RemoteAddr())
}

func (h *websocketHandle) RoomHandler(ws *websocket.Conn) {
	roomID := ws.Request().URL.Query().Get("id")
	username := ws.Request().URL.Query().Get("name")
	fmt.Printf("Room %s: %s has joined\n", roomID, username)

	if err := h.joinRoom(roomID, ws); err != nil {
		log.Println(err.Error())
		return
	}

	msg := fmt.Sprintf("%s has joined\n", username)
	h.broadcast(roomID, []byte(msg))

	h.read(roomID, ws)
}

func (h *websocketHandle) CreateRoomHandler(ws *websocket.Conn) {
	roomID := ws.Request().URL.Query().Get("id")

	if err := h.newRoom(roomID); err != nil {
		log.Println(err)
	}
}

func (h *websocketHandle) read(roomID string, ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Printf("Read error: %s", err.Error())
			continue
		}
		msg := buf[:n]
		h.broadcast(roomID, msg)
	}
}

func (h *websocketHandle) broadcast(roomID string, b []byte) {
	for ws := range h.rooms[roomID].conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				log.Printf("Broadcast error: %s", err.Error())
			}
		}(ws)
	}
}

func (h *websocketHandle) joinRoom(roomID string, ws *websocket.Conn) error {
	if _, ok := h.rooms[roomID]; !ok {
		msg := fmt.Sprintf("Room %s doesn't exists", roomID)
		if _, err := ws.Write([]byte(msg)); err != nil {
			log.Printf("Broadcast error: %s", err.Error())
		}
		return &exception.RoomDoesNotExists{RoomID: roomID}
	}

	h.rooms[roomID].conns[ws] = true
	msg := fmt.Sprintf("Welcome to room %s", roomID)
	ws.Write([]byte(msg))

	return nil
}

func (h *websocketHandle) newRoom(roomID string) error {
	if _, ok := h.rooms[roomID]; ok {
		return &exception.RoomAlreadyExists{RoomID: roomID}
	}

	// create room
	h.rooms[roomID] = connection{conns: make(map[*websocket.Conn]bool)}

	return nil
}

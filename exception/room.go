package exception

import "fmt"

type RoomAlreadyExists struct {
	RoomID string
}

type RoomDoesNotExists struct {
	RoomID string
}

func (e *RoomAlreadyExists) Error() string {
	return fmt.Sprintf("room id: %s already exists", e.RoomID)
}

func (e *RoomDoesNotExists) Error() string {
	return fmt.Sprintf("room id: %s doesn't exists", e.RoomID)
}

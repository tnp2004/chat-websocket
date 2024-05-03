package exception

import "fmt"

type ListeningFailed struct {
	Addr string
}

func (e *ListeningFailed) Error() string {
	return fmt.Sprintf("Listening on %s failed", e.Addr)
}

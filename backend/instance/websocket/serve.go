package websocket

import (
	"sync"

	"github.com/gofiber/websocket/v2"

	"jpc16-telnet-game/util/log"
	"jpc16-telnet-game/util/value"
)

var cnn = &webSocketConn{
	Conn:  nil,
	Mutex: new(sync.Mutex),
}

var InitialFunc func()

func Serve(conn *websocket.Conn) {
	HandleConnectionSwitch(cnn)

	// * Assign connection
	cnn.Mutex.Lock()
	cnn.Conn = conn
	cnn.Mutex.Unlock()

	// * Emit initial message
	InitialFunc()

	for {
		t, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if t != websocket.TextMessage {
			break
		}

		cnn.Emit(&OutboundMessage{
			Event:   EchoEvent,
			Payload: p,
		})
	}

	// * Close connection
	if err := conn.Close(); err != nil {
		log.Error("Unable to close connection", err)
	}

	// * Reset player connection
	cnn.Conn = nil

	// * Unlock in case of connection switch
	if value.MutexLocked(cnn.Mutex) {
		cnn.Mutex.Unlock()
	}
}

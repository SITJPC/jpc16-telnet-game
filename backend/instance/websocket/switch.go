package websocket

import (
	"jpc16-telnet-game/util/log"
)

func HandleConnectionSwitch(conn *webSocketConn) {
	// * Connection switch
	if conn.Conn != nil {
		log.Debug("Connection switch")
		conn.Emit(&OutboundMessage{
			Event:   ConnectionSwitchEvent,
			Payload: nil,
		})

		conn.Mutex.Lock()
		if err := conn.Conn.Close(); err != nil {
			log.Error("Unable to close connection", err)
		}
	}
}

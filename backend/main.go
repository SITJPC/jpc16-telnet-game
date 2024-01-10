package main

import (
	"jpc16-telnet-game/common/config"
	"jpc16-telnet-game/common/fiber"
	"jpc16-telnet-game/common/mng"
	"jpc16-telnet-game/instance/websocket"
)

func main() {
	config.Init()
	mng.Init()
	websocket.Init()
	fiber.Init()
}

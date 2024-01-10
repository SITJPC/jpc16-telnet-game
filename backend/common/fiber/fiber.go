package fiber

import (
	"github.com/gofiber/fiber/v2"

	cc "jpc16-telnet-game/common"
	"jpc16-telnet-game/common/fiber/middleware"
	"jpc16-telnet-game/endpoint"
	"jpc16-telnet-game/type/response"
	"jpc16-telnet-game/util/log"
	"jpc16-telnet-game/util/text"
)

func Init() {
	// Initialize fiber instance
	app := fiber.New(fiber.Config{
		AppName:       "JPC16 Core [" + text.Commit + "]",
		ErrorHandler:  ErrorHandler,
		Prefork:       false,
		StrictRouting: true,
	})

	// Register root endpoint
	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(response.Info("JPC16 API ROOT"))
	})

	// Register API endpoints
	apiGroup := app.Group("api/")
	apiGroup.Use(middleware.Recover())
	endpoint.Init(apiGroup)

	// Register websocket endpoints
	wsGroup := app.Group("ws/")
	InitWs(wsGroup)

	// Register not found endpoint
	app.Use(NotFoundHandler)

	// Startup
	err := app.Listen(*cc.Config.Address)
	if err != nil {
		log.Fatal("Unable to start fiber instance", err)
	}
}

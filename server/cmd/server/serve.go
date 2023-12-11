package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ebonian/where2kee/server/pkg/config"
	"github.com/ebonian/where2kee/server/pkg/middleware"
	"github.com/ebonian/where2kee/server/pkg/route"
	"github.com/ebonian/where2kee/server/platform/database"
	"github.com/ebonian/where2kee/server/platform/logger"
	"github.com/gofiber/fiber/v2"
)

func Serve() {
	appConfig := config.AppCfg()
	app := fiber.New(config.FiberConfig())

	log := logger.GetLogger()

	middleware.FiberMiddleware(app)

	route.GeneralRoute(app)
	route.BuildingRoute(app)
	route.ToiletRoute(app)
	route.NotFoundRoute(app)

	database.InitializeDB()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigCh
		log.Infoln("Shutting down server...")
		_ = app.Shutdown()
	}()

	serverAddr := fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)
	if err := app.Listen(serverAddr); err != nil {
		log.Errorf("Oops... server is not running! error: %v", err)
	}
}

package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"service/pkg/config"
	"service/pkg/handlers"
	"service/pkg/middleware"
	"service/pkg/routes"
)

type ServerManager struct{}

type Server interface {
	Start()
}

func NewServerManager() *ServerManager {
	return &ServerManager{}
}

func (sm *ServerManager) Start() {
	// Configuring the logger
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("Starting server...")

	// Initialize gin Engine
	router := gin.New()
	router.Use(middleware.LoggingMiddleware())
	gin.SetMode(gin.ReleaseMode)

	// Initialize config
	var configProvider config.Config = config.NewConfigManager()

	// Initialize handlers
	var handlersProvider handlers.Handlers = handlers.NewHandlersManager(
		configProvider.GetConfig().Database,
		configProvider.GetConfig().Enrich,
	)

	// Initialize routes
	var routesProvider routes.Routes = routes.NewRoutesManager(handlersProvider)
	routesProvider.Init(router)

	// Start server
	if err := router.Run(fmt.Sprintf(":%s", configProvider.GetConfig().App.Port)); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Server is stopped")
}
